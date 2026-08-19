package index

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"sync/atomic"
	"time"

	_ "github.com/whosonfirst/go-whosonfirst/v4/iterate/git"
	_ "github.com/whosonfirst/go-whosonfirst/v4/iterate/parquet"

	"github.com/sfomuseum/geocoder/coarse"
	"github.com/whosonfirst/go-whosonfirst/v4/feature/alt"
	"github.com/whosonfirst/go-whosonfirst/v4/uri"
)

func Run(ctx context.Context) error {

	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	opts, err := OptionsFromFlagSet(ctx, fs)

	if err != nil {
		return err
	}

	return RunWithOptions(ctx, opts)
}

func RunWithOptions(ctx context.Context, opts *Options) error {

	if opts.Verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	logger := slog.Default()

	if opts.Geocoder == nil {
		return fmt.Errorf("Missing geocoder")
	}

	defer opts.Geocoder.Close()

	t1 := time.Now()

	if opts.IndexJuggling {

		err := opts.Geocoder.PreIndex(ctx)

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

	for rec, err := range opts.Iterator.Iterate(ctx, opts.IteratorSources...) {

		if err != nil {
			return fmt.Errorf("Iterator yielded an error, %w", err)
		}

		new_count := atomic.AddInt64(&count, 1)

		if opts.Offset > 0 && new_count < opts.Offset {
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

		wof_opts := &coarse.NewWhosOnFirstRecordOptions{
			Body:     body,
			Embedder: opts.Embedder,
			EmbedderModels: []string{
				opts.EmbeddingsModel,
			},
		}

		rec, err := coarse.NewWhosOnFirstRecord(ctx, wof_opts)

		if err != nil {
			return fmt.Errorf("Failed to create new record, %w", err)
		}

		if opts.ExcludeNullIsland {

			if rec.Centroid.Lon() == 0 && rec.Centroid.Lat() == 0 {
				continue
			}
		}

		t1 := time.Now()

		if fresh {

			exists, has_changed, err := opts.Geocoder.HasRecordHashChanged(ctx, rec)

			if err != nil {
				return fmt.Errorf("Failed to determine if record hash has changed, %w", err)
			}

			// logger.Info("Has changed", "id", rec.Id, "changed", has_changed, "exists", exists)

			if !has_changed {
				continue
			}

			if exists && prune {

				err := opts.Geocoder.RemoveRecord(ctx, rec.Id)

				if err != nil {
					return fmt.Errorf("Failed to remove record, %w", err)
				}
			}
		}

		err = opts.Geocoder.AddRecord(ctx, rec)

		if err != nil {
			return fmt.Errorf("Failed to index records, %w", err)
		}

		go atomic.AddInt64(&tti, time.Since(t1).Milliseconds())
	}

	err := opts.Geocoder.Flush(ctx)

	if err != nil {
		return fmt.Errorf("Failed to flush database, %w", err)
	}

	logger.Info("Indexing complete", "seen", opts.Iterator.Seen(), "time", time.Since(t1), "average (ms)", avg_tti())

	if opts.IndexJuggling {

		t2 := time.Now()

		err = opts.Geocoder.PostIndex(ctx)

		if err != nil {
			return fmt.Errorf("Post-indexing failed, %w", err)
		}

		logger.Info("Post-indexing complete", "time", time.Since(t2), "time (total)", time.Since(t1))
	}

	return nil
}
