package main

import (
	"context"
	"log"

	"github.com/sfomuseum/geocoder/app/coarse/server"
)

func main() {

	ctx := context.Background()
	err := server.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to server geocoder, %v", err)
	}
}
