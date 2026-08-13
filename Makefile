CWD=$(shell pwd)

GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

vuln:
	govulncheck -show verbose ./...

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-coarse-geocoder-index cmd/wof-coarse-geocoder-index/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-coarse-geocoder-index-tgn cmd/wof-coarse-geocoder-index-tgn/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-coarse-geocoder-query cmd/wof-coarse-geocoder-query/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-coarse-geocoder-query-fs cmd/wof-coarse-geocoder-query-fs/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-coarse-geocoder-server cmd/wof-coarse-geocoder-server/main.go

# This does not compile because modernc.org/libc doesn not target GOOS=js
wasmjs:
	GOOS=js GOARCH=wasm \
		go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags wasmjs \
		-o work/geocoder-query.wasm \
		cmd/wof-coarse-geocoder-query-wasm/main.go

# This does not compile because modernc.org/libc doesn not target GOOS=wasip1
wasip1:
	GOARCH=wasm GOOS=wasip1 \
		go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags wasip1 \
		-o work/query-p1.wasm \
		./cmd/wof-coarse-geocoder-query-wasi/main.go

# I don't know if the compiles because tinygo doesn't get that far failing on unsupported
# net/http/httputil imported by vendor deps...
wasip2:
	tinygo build -target wasip2 -o work/query-p2.wasm ./cmd/wof-coarse-geocoder-query-wasi/main.go

lambda:
	@make lambda-server-fs

lambda-server-fs:
	if test -f bootstrap; then rm -f bootstrap; fi
	if test -f geocoder-server.zip; then rm -f geocoder-server.zip; fi
	GOARCH=arm64 GOOS=linux go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags lambda.norpc -o bootstrap cmd/wof-coarse-geocoder-server-fs/main.go
	zip geocoder-server.zip bootstrap
	rm -f bootstrap

server:
	go run -mod $(GOMOD) cmd/wof-coarse-geocoder-server/main.go \
		-demo \
		-verbose \
		-server-uri http://localhost:8080 \
		-geocoder-uri $(GEOCODER_URI)
