CWD=$(shell pwd)

GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

vuln:
	govulncheck -show verbose ./...

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-coarse-geocoder-index cmd/wof-coarse-geocoder-index/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-coarse-geocoder-query cmd/wof-coarse-geocoder-query/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-coarse-geocoder-server cmd/wof-coarse-geocoder-server/main.go

server:
	go run -mod $(GOMOD) cmd/wof-coarse-geocoder-server/main.go \
		-demo \
		-verbose \
		-server-uri http://localhost:8080 \
		-geocoder-uri $(GEOCODER_URI)
