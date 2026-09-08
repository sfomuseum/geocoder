package parquet

import (
	"errors"
	"io"
	"sync"
)

var _ ReadCloserAt = (*cachedReaderAt)(nil)

// Create a wrapper interface that bundles all required behaviors
type ReadCloserAt interface {
	io.ReaderAt
	io.Closer
}

// Ensure cachedReaderAt implements our new interface
var _ ReadCloserAt = (*cachedReaderAt)(nil)

// Internal data structure to wrap net/http.Response.Body instances
// so we don't read the whole document over the wire before processing
type cachedReaderAt struct {
	src    io.ReadCloser
	buf    []byte
	closed bool
	err    error
	mu     sync.Mutex
}

func NewCachedReaderAt(src io.ReadCloser) ReadCloserAt {
	return &cachedReaderAt{
		src: src,
	}
}

func (c *cachedReaderAt) ReadAt(p []byte, off int64) (int, error) {

	if off < 0 {
		return 0, errors.New("cachedReaderAt.ReadAt: negative offset")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	reqEnd := off + int64(len(p))

	// Reuse a single buffer allocation across iterations for this call
	var tmp []byte

	// Pull data into cache until we satisfy the request or hit EOF/error
	for reqEnd > int64(len(c.buf)) && !c.closed {

		if tmp == nil {
			tmp = make([]byte, 4096)
		}

		n, err := c.src.Read(tmp)
		if n > 0 {
			c.buf = append(c.buf, tmp[:n]...)
		}
		if err != nil {
			c.closed = true
			if err != io.EOF {
				c.err = err // Store the unexpected network error
			}
			break
		}
	}

	bufLen := int64(len(c.buf))

	// Case 1: The requested offset starts past the total data we have
	if off >= bufLen {
		if c.err != nil {
			return 0, c.err
		}
		return 0, io.EOF
	}

	// Case 2: We have partial data matching the offset, but less than requested
	if reqEnd > bufLen {
		n := copy(p, c.buf[off:])
		if c.err != nil {
			return n, c.err
		}
		return n, io.EOF
	}

	// Case 3: We have a full match in the cache
	n := copy(p, c.buf[off:reqEnd])
	return n, nil
}

// Close safely tears down the underlying reader
func (c *cachedReaderAt) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Prevent double-closing errors
	if c.closed && c.err == io.ErrClosedPipe {
		return nil
	}

	c.closed = true
	c.err = io.ErrClosedPipe

	return c.src.Close()
}
