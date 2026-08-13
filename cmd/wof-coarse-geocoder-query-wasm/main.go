//go:build wasmjs
package main

// This code will NOT compile. It will not compile because it depends
// on the `modernc.org/sqlite/vfs` package to bunlde a SQLite database in
// an embedded filesystem which in turn depends on the `modernc.org/libc`
// package which does not target the "JS" operating system. It is included
// for documentary purposes and in the hope that "JS" will be a build option.

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
