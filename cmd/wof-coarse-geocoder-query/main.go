package main

import (
	"context"
	"log"

	"github.com/sfomuseum/geocoder/app/coarse/query"
)

func main() {

	ctx := context.Background()
	err := query.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to query geocoder, %v", err)
	}
}
