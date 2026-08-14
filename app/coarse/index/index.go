package index

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/url"
	"slices"
	"sync/atomic"
	"time"

	_ "github.com/whosonfirst/go-whosonfirst/v4/iterate/git"
	_ "github.com/whosonfirst/go-whosonfirst/v4/iterate/parquet"

	"github.com/aaronland/go-json-query"
	"github.com/sfomuseum/geocoder/coarse"
	"github.com/sfomuseum/go-embeddings"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst/v4/feature/alt"
	"github.com/whosonfirst/go-whosonfirst/v4/iterate"
	"github.com/whosonfirst/go-whosonfirst/v4/uri"
)

func Run(ctx context.Context) error {

	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	flagset.Parse(fs)

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	logger := slog.Default()

	gc, err := coarse.NewSQLGeocoder(ctx, geocoder_uri)

	if err != nil {
		return fmt.Errorf("Failed to create geocoder, %w", err)
	}

	defer gc.Close()

	if exclude_deprecated || exclude_superseded || exclude_funky {

		exclude_deprecated_path := "properties.edtf:deprecated=.*"
		exclude_superseded_path := "properties.wof:superseded_by=.*"
		exclude_funky_path := "propertiees.mz:is_funky=1"

		u, err := url.Parse(iterator_uri)

		if err != nil {
			return fmt.Errorf("Failed to parse iterator URI, %w", err)
		}

		q := u.Query()

		to_exclude := make([]string, 0)

		if exclude_deprecated {
			to_exclude = append(to_exclude, exclude_deprecated_path)
		}

		if exclude_superseded {
			to_exclude = append(to_exclude, exclude_superseded_path)
		}

		if exclude_funky {
			to_exclude = append(to_exclude, exclude_funky_path)
		}

		_, ok := q["exclude"]

		if ok {

			for _, v := range q["exclude"] {

				if !slices.Contains(to_exclude, v) {
					to_exclude = append(to_exclude, v)
				}
			}
		}

		q["exclude"] = to_exclude

		q.Set("exclude_mode", query.QUERYSET_MODE_ANY)
		u.RawQuery = q.Encode()
		iterator_uri = u.String()

		logger.Info("Rewrote iterator URI", "uri", iterator_uri)
	}

	iter, err := iterate.NewIterator(ctx, iterator_uri)

	if err != nil {
		return fmt.Errorf("Failed to create new iteratr, %w", err)
	}

	t1 := time.Now()

	if index_juggling {

		err = gc.PreIndex(ctx)

		if err != nil {
			return fmt.Errorf("Pre-indexing failed, %w", err)
		}

		logger.Info("Pre-indexing complete", "time", time.Since(t1))
	}

	count := int64(0)
	tti := int64(0)

	ticker := time.NewTicker(60 * time.Second)
	done_ch := make(chan bool)

	defer func() {
		done_ch <- true
		ticker.Stop()
	}()

	avg_tti := func() float64 {
		current_tti := atomic.LoadInt64(&tti)
		current_count := atomic.LoadInt64(&count)
		return float64(current_tti) / float64(current_count)
	}

	go func() {

		for {
			select {
			case <-done_ch:
				return
			case <-ticker.C:
				current_count := atomic.LoadInt64(&count)
				logger.Info("Indexing stats", "elapsed", time.Since(t1), "seen", current_count, "average (ms)", avg_tti())
			}
		}
	}()

	embedder, err := embeddings.NewEmbedder32(ctx, "ollama://?model=embeddingsgemma")

	if err != nil {
		return fmt.Errorf("Failed to create embedder, %w", err)
	}

	iterator_uris := fs.Args()

	for rec, err := range iter.Iterate(ctx, iterator_uris...) {

		if err != nil {
			return fmt.Errorf("Iterator yielded an error, %w", err)
		}

		new_count := atomic.AddInt64(&count, 1)

		if offset > 0 && new_count < offset {
			rec.Body.Close()
			continue
		}

		is_alt, err := uri.IsAlternateGeometry(rec.Path)

		if err != nil {
			logger.Debug("Failed to determine if alt", "path", rec.Path, "error", err)
		}

		if is_alt {
			rec.Body.Close()
			continue
		}

		body, err := rec.ReadAllAndClose()

		if err != nil {
			return fmt.Errorf("Failed to read body, %w", err)
		}

		if alt.IsAlt(body) {
			continue
		}

		opts := &coarse.NewWhosOnFirstRecordOptions{
			Body:     body,
			Embedder: embedder,
		}

		rec, err := coarse.NewWhosOnFirstRecord(ctx, opts)

		if err != nil {
			return fmt.Errorf("Failed to create new record, %w", err)
		}

		if exclude_nullisland {

			if rec.Centroid.Lon() == 0 && rec.Centroid.Lat() == 0 {
				continue
			}
		}

		t1 := time.Now()

		if fresh {

			exists, has_changed, err := gc.HasRecordHashChanged(ctx, rec)

			if err != nil {
				return fmt.Errorf("Failed to determine if record hash has changed, %w", err)
			}

			// logger.Info("Has changed", "id", rec.Id, "changed", has_changed, "exists", exists)

			if !has_changed {
				continue
			}

			if exists && prune {

				err := gc.RemoveRecord(ctx, rec.Id)

				if err != nil {
					return fmt.Errorf("Failed to remove record, %w", err)
				}
			}
		}

		err = gc.AddRecord(ctx, rec)

		if err != nil {
			return fmt.Errorf("Failed to index records, %w", err)
		}

		go atomic.AddInt64(&tti, time.Since(t1).Milliseconds())
	}

	err = gc.Flush(ctx)

	if err != nil {
		return fmt.Errorf("Failed to flush database, %w", err)
	}

	logger.Info("Indexing complete", "seen", iter.Seen(), "time", time.Since(t1), "average (ms)", avg_tti())

	if index_juggling {

		t2 := time.Now()

		err = gc.PostIndex(ctx)

		if err != nil {
			return fmt.Errorf("Post-indexing failed, %w", err)
		}

		logger.Info("Post-indexing complete", "time", time.Since(t2), "time (total)", time.Since(t1))
	}

	return nil
}
