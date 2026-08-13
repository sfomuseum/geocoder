//go:build wasmjs
package main

import (
	"context"
	"fmt"
	"log"
	"syscall/js"
	
	"github.com/sfomuseum/geocoder/x/wasm"	
)

func main() {

	ctx := context.Background()

	query_func, err := wasm.NewQueryFuncWithFS(ctx)

	if err != nil {
		log.Fatalf("Failed to create query func, %v", err)
	}
	
	defer query_func.Release()

	js.Global().Set("geocoder_query", query_func)

	c := make(chan struct{}, 0)

	log.Println("WASM geocoder query function initialized")
	<-c
	
}
