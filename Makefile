CWD=$(shell pwd)

GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

vuln:
	govulncheck -show verbose ./...

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-coarse-geocoder-index cmd/wof-coarse-geocoder-index/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-coarse-geocoder-index-tgn cmd/wof-coarse-geocoder-index-tgn/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-coarse-geocoder-query cmd/wof-coarse-geocoder-query/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-coarse-geocoder-server cmd/wof-coarse-geocoder-server/main.go

lambda:
	@make lambda-server

lambda-server:
	if test -f bootstrap; then rm -f bootstrap; fi
	if test -f geocoder-server.zip; then rm -f geocoder-server.zip; fi
	GOARCH=arm64 GOOS=linux go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags lambda.norpc -o bootstrap cmd/wof-coarse-geocoder-server/main.go
	zip geocoder-server.zip bootstrap
	rm -f bootstrap

server:
	go run -mod $(GOMOD) cmd/wof-coarse-geocoder-server/main.go \
		-demo \
		-verbose \
		-server-uri http://localhost:8080 \
		-geocoder-uri $(GEOCODER_URI)
