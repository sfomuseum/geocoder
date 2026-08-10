package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/url"
	"os"
	"slices"
	"sync/atomic"
	"time"

	_ "github.com/whosonfirst/go-whosonfirst/v4/iterate/git"
	_ "github.com/whosonfirst/go-whosonfirst/v4/iterate/parquet"

	"github.com/aaronland/go-json-query"
	"github.com/sfomuseum/geocoder/coarse"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst/v4/feature/alt"
	"github.com/whosonfirst/go-whosonfirst/v4/iterate"
	"github.com/whosonfirst/go-whosonfirst/v4/uri"
)

func main() {

	var iterator_uri string
	var offset int64
	var geocoder_uri string
	var fresh bool
	var prune bool
	var index_juggling bool

	var exclude_deprecated bool
	var exclude_superseded bool
	var exclude_funky bool
	var exclude_nullisland bool

	var verbose bool

	fs := flagset.NewFlagSet("index")

	fs.StringVar(&iterator_uri, "iterator-uri", "repo://", "A registered whosonfirst/go-whosonfirst/v4/iterate.Iterate URI.")
	fs.StringVar(&geocoder_uri, "geocoder-uri", "sql://sqlite?dsn=:memory:", "A registered sfomuseum/geocoder/coarse.Geocoder URI.")
	fs.Int64Var(&offset, "offset", 0, "Optional document offset to start indexing from.")
	fs.BoolVar(&fresh, "fresh", false, "This flags signals that a fresh database is being indexed disabling checks for existing or updated records.")
	fs.BoolVar(&prune, "prune", false, "Prune existing records before (re)adding them to the database.")
	fs.BoolVar(&index_juggling, "index-juggling", true, "Perform indexing speed optiomizations. This will include dropping existing indices and the FTS table prior to indexing and (re)adding them at the end.")

	fs.BoolVar(&exclude_deprecated, "exclude-deprecated", true, "Do not index records which have been deprecated.")
	fs.BoolVar(&exclude_superseded, "exclude-superseded", true, "Do not index records which have been superseded.")
	fs.BoolVar(&exclude_funky, "exclude-funky", true, "Do not index records which have been flagged as \"funky\".")
	fs.BoolVar(&exclude_nullisland, "exclude-nullisland", true, "Do not index records that are \"visiting\" Null Island (have 0,0 coordinate data).")

	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Index one or more Who's On First data sources in a (coarse) geocoding database.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s [options] uri(N) uri(N) uri(N)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		fs.PrintDefaults()
	}

	flagset.Parse(fs)

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	logger := slog.Default()

	ctx := context.Background()

	gc, err := coarse.NewSQLGeocoder(ctx, geocoder_uri)

	if err != nil {
		log.Fatalf("Failed to create geocoder, %v", err)
	}

	defer gc.Close()

	if exclude_deprecated || exclude_superseded || exclude_funky {

		exclude_deprecated_path := "properties.edtf:deprecated=.*"
		exclude_superseded_path := "properties.wof:superseded_by=.*"
		exclude_funky_path := "propertiees.mz:is_funky=1"

		u, err := url.Parse(iterator_uri)

		if err != nil {
			log.Fatalf("Failed to parse iterator URI, %v", err)
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
		log.Fatalf("Failed to create new iteratr, %v", err)
	}

	t1 := time.Now()

	if index_juggling {

		err = gc.PreIndex(ctx)

		if err != nil {
			log.Fatalf("Pre-indexing failed, %v", err)
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

	iterator_uris := fs.Args()

	for rec, err := range iter.Iterate(ctx, iterator_uris...) {

		if err != nil {
			log.Fatalf("Iterator yielded an error, %v", err)
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
			log.Fatalf("Failed to read body, %v", err)
		}

		if alt.IsAlt(body) {
			continue
		}

		rec, err := coarse.NewWhosOnFirstRecord(ctx, body)

		if err != nil {
			log.Fatalf("Failed to create new record, %v", err)
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
				log.Fatalf("Failed to determine if record hash has changed, %v", err)
			}

			// logger.Info("Has changed", "id", rec.Id, "changed", has_changed, "exists", exists)

			if !has_changed {
				continue
			}

			if exists && prune {

				err := gc.RemoveRecord(ctx, rec.Id)

				if err != nil {
					log.Fatalf("Failed to remove record, %v", err)
				}
			}
		}

		err = gc.AddRecord(ctx, rec)

		if err != nil {
			log.Fatalf("Failed to index records, %v", err)
		}

		go atomic.AddInt64(&tti, time.Since(t1).Milliseconds())
	}

	err = gc.Flush(ctx)

	if err != nil {
		log.Fatalf("Failed to flush database, %v", err)
	}

	logger.Info("Indexing complete", "seen", iter.Seen(), "time", time.Since(t1), "average (ms)", avg_tti())

	if index_juggling {

		t2 := time.Now()

		err = gc.PostIndex(ctx)

		if err != nil {
			log.Fatalf("Post-indexing failed, %v", err)
		}

		logger.Info("Post-indexing complete", "time", time.Since(t2), "time (total)", time.Since(t1))
	}

}
