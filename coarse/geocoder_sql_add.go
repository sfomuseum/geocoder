package coarse

import (
	"context"
	db_sql "database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"slices"
	"strings"
	_ "sync/atomic"
	"time"

	_ "modernc.org/sqlite"
	_ "modernc.org/sqlite/vec"

	geocoder_sql "github.com/sfomuseum/geocoder/coarse/sql"
)

func (g *SQLGeocoder) AddRecord(ctx context.Context, rec *Record) error {

	g.mu.Lock()
	defer g.mu.Unlock()

	g.records = append(g.records, rec)

	if len(g.records) >= g.batch_size {

		err := g.addRecords(ctx, g.records...)

		if err != nil {
			return err
		}

		g.records = make([]*Record, 0)
	}

	return nil
}

func (g *SQLGeocoder) addRecords(ctx context.Context, records ...*Record) error {

	logger := slog.Default()
	logger.Debug("Add bulk records", "count", len(records), "workers", g.bulk_workers)

	tx, err := g.db.BeginTx(ctx, &db_sql.TxOptions{
		Isolation: db_sql.LevelDefault,
		ReadOnly:  false,
	})

	if err != nil {
		return fmt.Errorf("Failed to begin transaction, %w", err)
	}

	t1 := time.Now()

	defer func() {

		tx.Rollback()

		if err != nil && err != db_sql.ErrTxDone {
			logger.Error("Failed to rollback transaction", "error", err)
		}

		logger.Debug("Time to bulk index records", "count", len(records), "time", time.Since(t1))
	}()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	done_ch := make(chan bool)
	err_ch := make(chan error)

	throttle := make(chan bool, g.bulk_workers)

	for i := 0; i < g.bulk_workers; i++ {
		throttle <- true
	}

	for _, rec := range records {

		<-throttle

		select {
		case <-ctx.Done():
			break
		default:
			// pass
		}

		go func(rec *Record) {

			logger := slog.Default()
			logger = logger.With("id", rec.Id)

			defer func() {
				throttle <- true
				done_ch <- true
			}()

			//

			logger.Info("STORE", "id", rec.Id)
			
			record_id, err := g.storeIdentifier(ctx, tx, rec.Id)

			if err != nil {
				err_ch <- err
				return
			}

			logger = logger.With("record id", record_id)

			parent_id, err := g.storeIdentifier(ctx, tx, rec.ParentId)

			if err != nil {
				err_ch <- err
				return
			}

			others := make([]string, 0)

			for _, h := range rec.Hierarchies {

				for _, v := range h {

					if !slices.Contains(others, v) {
						others = append(others, v)
					}
				}
			}

			for _, other_id := range others {
				_, err := g.storeIdentifier(ctx, tx, other_id)

				if err != nil {
					logger.Error("Failed to store identifier for other ID", "other", other_id, "error", err)
				}
			}

			//

			enc_hierarchies, err := json.Marshal(rec.Hierarchies)

			if err != nil {
				err_ch <- fmt.Errorf("Failed to marshal hierarchies, %w", err)
				return
			}

			record_hash, err := rec.Hash()

			if err != nil {
				err_ch <- fmt.Errorf("Failed to hash record, %w", err)
				return
			}

			rec_q := fmt.Sprintf("INSERT INTO %s (id, parent_id, name, placetype, latitude, longitude, country, inception, cessation, hierarchies, is_current, population_rank, record_hash) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", g.tableName("records"))

			_, err = tx.ExecContext(ctx, rec_q, record_id, parent_id, rec.Name, rec.Placetype, rec.Centroid.Lat(), rec.Centroid.Lon(), rec.Country, rec.Inception, rec.Cessation, string(enc_hierarchies), rec.IsCurrent, rec.PopulationRank, record_hash)

			if err != nil {
				logger.Error("Failed to create record row", "error", err)
				err_ch <- err
				return
			}

			// logger.Info("WTF")
			// Placetypes (alt)

			pt_stq := fmt.Sprintf("INSERT INTO %s (record_id, placetype) VALUES(?, ?)", g.tableName("placetypes_alt"))

			pt_st, err := tx.Prepare(pt_stq)

			if err != nil {
				err_ch <- fmt.Errorf("Failed to prepare placetypes statement, %w", err)
				return
			}

			defer pt_st.Close()

			for _, pt := range rec.PlacetypeAlt {

				_, err = pt_st.ExecContext(ctx, record_id, pt)

				if err != nil {
					err_ch <- fmt.Errorf("Failed to execute placetype statement, %w", err)
					return
				}
			}

			// Ancestors

			anc_stq := fmt.Sprintf("INSERT INTO %s (record_id, ancestor_id) VALUES(?, ?)", g.tableName("ancestors"))
			anc_st, err := tx.Prepare(anc_stq)

			if err != nil {
				err_ch <- fmt.Errorf("Failed to prepare ancestors statement, %w", err)
				return
			}

			defer anc_st.Close()

			ancestors := make([]int64, 0)

			for _, hier := range rec.Hierarchies {

				for _, str_id := range hier {

					id, err := g.retrieveIdentifier(ctx, tx, str_id)

					if err != nil {
						logger.Error("Failed to retrieve identifier", "id", str_id, "error", err)
						continue
					}

					if !slices.Contains(ancestors, id) {
						ancestors = append(ancestors, id)
					}
				}
			}

			for _, id := range ancestors {

				_, err = anc_st.ExecContext(ctx, record_id, id)

				if err != nil {
					err_ch <- fmt.Errorf("Failed to execute ancestors statement, %w", err)
					return
				}
			}

			// Bounds

			bounds_stq := fmt.Sprintf("INSERT INTO %s (minx, miny, maxx, maxy, record_id) VALUES(?, ?, ?, ?, ?)", g.tableName("bounds"))
			bounds_st, err := tx.Prepare(bounds_stq)

			if err != nil {
				err_ch <- fmt.Errorf("Failed to prepare bounds statement, %w", err)
				return
			}

			defer bounds_st.Close()

			for _, b := range rec.Bounds {

				_, err = bounds_st.ExecContext(ctx, b.Min.X(), b.Min.Y(), b.Max.X(), b.Max.Y(), record_id)

				if err != nil {
					err_ch <- fmt.Errorf("Failed to execute bounds statement, %w", err)
					return
				}
			}

			// Dates

			start_outer, start_inner, end_inner, end_outer := rec.DateRanges()

			date_fields := []string{
				"record_id",
			}

			date_args := []any{
				record_id,
			}

			if start_outer != nil {
				date_fields = append(date_fields, "start_outer")
				date_args = append(date_args, start_outer.Unix())
			}

			if start_inner != nil {
				date_fields = append(date_fields, "start_inner")
				date_args = append(date_args, start_inner.Unix())
			}

			if end_inner != nil {
				date_fields = append(date_fields, "end_inner")
				date_args = append(date_args, end_inner.Unix())
			}

			if end_outer != nil {
				date_fields = append(date_fields, "end_outer")
				date_args = append(date_args, end_outer.Unix())
			}

			if len(date_fields) > 1 {

				date_placeholders := make([]string, len(date_fields))

				for i, _ := range date_fields {
					date_placeholders[i] = "?"
				}

				date_q := fmt.Sprintf("INSERT OR REPLACE INTO %s (%s) VALUES (%s)", g.tableName("dates"), strings.Join(date_fields, ","), strings.Join(date_placeholders, ","))

				_, err := tx.ExecContext(ctx, date_q, date_args...)

				if err != nil {
					err_ch <- fmt.Errorf("Failed to insert dates, %v", err)
					return
				}
			}

			// Tokens

			tok_stq := fmt.Sprintf("INSERT INTO %s (record_id, token, lang, tag) VALUES(?, ?, ?, ?)", g.tableName("tokens"))
			tok_st, err := tx.Prepare(tok_stq)

			if err != nil {
				err_ch <- fmt.Errorf("Failed to prepare token statement, %w", err)
				return
			}

			defer tok_st.Close()

			for lang, tag_tokens := range rec.Tokens {

				for tag, tokens := range tag_tokens {
					_, err = tok_st.ExecContext(ctx, record_id, strings.Join(tokens, " "), lang, tag)

					if err != nil {
						err_ch <- fmt.Errorf("Failed to execute token statement, %w", err)
						return
					}
				}
			}

			// Vectors
			// To do: Determine if buffering vector embeddings insertion _outside_ of any given record
			// and doing those updates in batches with entirely separate transactions will speed up overall
			// indexing time. This will require logic to ensure buffers get flushed and play nicely with
			// per-record transactions.

			if rec.VectorEmbeddings != nil && len(rec.VectorEmbeddings) > 0 {

				slog.Debug("Add vector embeddings", "id", rec.Id, "count", len(rec.VectorEmbeddings))

				emb_table := g.tableName("embeddings")

				var vec_q string

				switch g.vector_compression {
				case geocoder_sql.SQLiteVecQuantizeCompression:
					vec_q = fmt.Sprintf("INSERT OR REPLACE INTO %s (rowid, embedding) VALUES (?, vec_quantize_binary(?))", emb_table)
				case geocoder_sql.SQLiteVecMatroyshkaCompression:
					vec_q = fmt.Sprintf("INSERT OR REPLACE INTO %s (rowid, embedding) VALUES (?, vec_normalize(vec_slice(?, 0, %d)))", emb_table, geocoder_sql.SQLiteMatroyshkaDimensions)
				case geocoder_sql.SQLiteVecDefaultCompression:
					vec_q = fmt.Sprintf("INSERT OR REPLACE INTO %s (rowid, embedding) VALUES (:id, vec_f32(:vector))", emb_table)
				default:
					err_ch <- fmt.Errorf("Invalid or unsupported compression, '%s'", g.vector_compression)
					return
				}

				vec_st, err := tx.Prepare(vec_q)

				if err != nil {
					err_ch <- fmt.Errorf("Failed to prepare vector statement, %w", err)
					return
				}

				defer vec_st.Close()

				vrec_q := fmt.Sprintf("INSERT OR REPLACE INTO %s (id, record_id, model, language, tag) VALUES (?, ?, ?, ?, ?)", g.tableName("embeddings_records"))

				vrec_st, err := tx.Prepare(vrec_q)

				if err != nil {
					err_ch <- fmt.Errorf("Failed to prepare vector record statement, %w", err)
					return
				}

				defer vrec_st.Close()

				del_q := fmt.Sprintf("DELETE FROM %s WHERE rowid = ?", emb_table)

				for _, v := range rec.VectorEmbeddings {

					for _, e := range v.Embeddings {

						enc_e, err := json.Marshal(e.Embeddings)

						if err != nil {
							err_ch <- fmt.Errorf("Failed to marshal embeddings, %w", err)
							return
						}

						vrec_id, err := g.uidForVectorRecord(ctx, tx, record_id, v.Model, e.Language, e.Tag)

						if err != nil {
							err_ch <- fmt.Errorf("Failed to derive UID for vector record, %w", err)
							return
						}

						_, err = tx.ExecContext(ctx, del_q, vrec_id)

						if err != nil {
							logger.Error("NOPE vec delet", "q", del_q, "id", vrec_id, "error", err)
							err_ch <- fmt.Errorf("Failed to delete row (%d) from mebeddings, %w", vrec_id, err)
							return
						}

						_, err = vec_st.ExecContext(ctx, vec_q, db_sql.Named("id", vrec_id), db_sql.Named("vector", string(enc_e)))

						if err != nil {
							logger.Error("NOPE vec", "q", vec_q, "error", err)
							err_ch <- fmt.Errorf("Failed to add embeddings, %w", err)
							return
						}

						_, err = vrec_st.ExecContext(ctx, vrec_id, record_id, v.Model, e.Language, e.Tag)

						if err != nil {
							logger.Error("NOPE vec r", "error", err)
							err_ch <- fmt.Errorf("Failed to add vector record row, %w", err)
							return
						}
					}
				}
			}

			// All done...

		}(rec)
	}

	remaining := len(records)

	for remaining > 0 {
		select {
		case <-done_ch:
			remaining -= 1
			// logger.Info("Bulk add", "remaining", remaining)
		case err := <-err_ch:
			return err
		}
	}

	err = tx.Commit()

	if err != nil {
		return fmt.Errorf("Failed to commit transaction, %w", err)
	}

	return nil
}
