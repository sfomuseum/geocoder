package main

import (
	"context"
	"log"

	"github.com/sfomuseum/geocoder/app/coarse/index"
)

func main() {

	ctx := context.Background()
	err := index.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to index geocoder, %v", err)
	}
}
