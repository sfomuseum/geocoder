//go:build wasmjs
package wasm

// This code will NOT compile. It will not compile because it depends
// on the `modernc.org/sqlite/vfs` package to bunlde a SQLite database in
// an embedded filesystem which in turn depends on the `modernc.org/libc`
// package which does not target the "JS" operating system. It is included
// for documentary purposes and in the hope that "JS" will be a build option.

import (
	"context"
	"encoding/json"
	"strings"
	"syscall/js"
	"database/sql"
	
	"github.com/sfomuseum/geocoder/coarse"
	"github.com/aaronland/go-pagination/countable"	
)

func NewQueryFuncWithFS(ctx context.Context) (js.Func, error) {

	var db *sql.DB

	gc_opts := coarse.NewSQLGeocoderOptions{
		Database: db,
	}

	gc, err := coarse.NewSQLGeocoder(ctx, gc_opts)

	if err != nil {
		return nil, fmt.Errorf("Failed to create geocoder, %v", err)
	}

	return QueryFunc(gc), nil
	
}

func QueryFunc(db coarse.Geocoder) js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		req_str := args[0].String()
		req_r := strings.NewReader(req_str)
		
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			resolve := args[0]
			reject := args[1]

			ctx := context.Background()
			
			var query_req *coarse.QueryRequest

			dec := json.NewDecoder(r)
			err := dec.Decode(&query_req)

			if err != nil {
				reject.Invoke(fmt.Sprintf("Failed to decode query request, %w", err))
				return
			}

			pg_opts, err := countable.NewCountableOptions()			

			if err != nil {
				reject.Invoke(fmt.Sprintf("Failed to create pagination options, %w", err))
				return
			}
			
			rsp, _, err := db.Query(ctx, query_req, pg_opts)

			if err != nil {
				reject.Invoke(fmt.Sprintf("Failed to decode query database, %w", err))
				return
			}

			enc_rsp, err := json.Marshal(rsp)

			if err != nil {
				reject.Invoke(fmt.Sprintf("Failed to encode query results, %w", err))
				return
			}

			resolve.Invoke(string(enc_rsp))
			return
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)		
	}
}
