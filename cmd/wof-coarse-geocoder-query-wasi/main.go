//go:build wasip1
package main

import (
	"context"
	"fmt"
	"encoding/json"
	
	"github.com/sfomuseum/geocoder/coarse"
	x_fs "github.com/sfomuseum/geocoder/x/fs"
)

func main() {

	args := strings.Join(flag.Args(), " ")
	fmt.Println(query(args))
}

func query(args string) string {

	ctx := context.Background()

	db_fs := x_fs.FS
	db_name := "sfomuseum_arch.db"

	gc, err := coarse.NewSQLGeocoderWithFS(ctx, db_fs, db_name)

	if err != nil {
		return fmt.Errorf("Failed to create geocoder, %w", err)
	}

	var req *coarse.QueryRequest

	err = json.Unmarshal([]byte(args), &req)

	if err != nil {
		return err.Error()
	}
	
	pg_opts, err := countable.NewCountableOptions()

	if err != nil {
		return err.Error()
	}

	rsp, _, err := gc.Query(ctx, req, pg_opts)

	if err != nil {
		return err.Error()
	}

	enc_rsp, err := json.Marshal(rsp)

	if err != nil {
		return err.Error()
	}

	return string(enc_rsp)
}
