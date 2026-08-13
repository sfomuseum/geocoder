// gobuild: vfs_remote
package vfs

import (
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
)

// RemoteHTTPFS implements Go's fs.FS interface
type RemoteHTTPFS struct {
	BaseURL string
	Client  *http.Client
}

func (f *RemoteHTTPFS) Open(name string) (fs.File, error) {

	u, err := url.JoinPath(f.BaseURL, name)

	req, err := http.NewRequest("HEAD", u, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to create HEAD request: %w", err)
	}

	rsp, err := f.Client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to remote server: %w", err)
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("remote file not found or inaccessible, HTTP status: %d", rsp.StatusCode)
	}

	str_length := rsp.Header.Get("Content-Length")

	if str_length == "" {
		return nil, fmt.Errorf("remote server failed to provide Content-Length header")
	}

	size, err := strconv.ParseInt(str_length, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("invalid Content-Length header format: %w", err)
	}

	if rsp.Header.Get("Accept-Ranges") == "none" {
		slog.Warn("Server explicitly explicitly states Accept-Ranges: none. Range reads may fail.")
	}

	return &HTTPFile{
		URL:      u,
		Client:   f.Client,
		fileSize: size,
		fileName: name,
	}, nil
}
