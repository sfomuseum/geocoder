package parquet

import (
	"net/http"
	"net/url"
	"os"
)

func OpenURI(uri string) (ReadCloserAt, int64, error) {

	var r ReadCloserAt
	var sz int64

	u, err := url.Parse(uri)

	if err != nil {
		return nil, 0, err
	}

	switch u.Scheme {
	case "http", "https":

		rsp, err := http.Get(uri)

		if err != nil {
			return nil, 0, err
		}

		r = NewCachedReaderAt(rsp.Body)
		sz = rsp.ContentLength

	default:

		f, err := os.Open(uri)

		if err != nil {
			return nil, 0, err
		}

		info, err := f.Stat()

		if err != nil {
			return nil, 0, err
		}

		r = f
		sz = info.Size()
	}

	return r, sz, nil
}
