package vfs

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"time"
)

// HTTPFile implements fs.File, io.ReaderAt, and io.Seeker to satisfy modernc.org/sqlite
type HTTPFile struct {
	URL      string
	Client   *http.Client
	fileSize int64
	fileName string
	offset   int64 // Tracking stateful pointer position
}

// Seek handles cursor movement requests issued by modernc's VFS logic
func (f *HTTPFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.offset = offset
	case io.SeekCurrent:
		f.offset += offset
	case io.SeekEnd:
		f.offset = f.fileSize + offset
	default:
		return 0, errors.New("invalid whence value")
	}

	if f.offset < 0 {
		return 0, errors.New("negative position error")
	}
	return f.offset, nil
}

// Read handles internal sequential block streaming
func (f *HTTPFile) Read(p []byte) (int, error) {

	if f.offset >= f.fileSize {
		return 0, io.EOF
	}

	n, err := f.ReadAt(p, f.offset)
	f.offset += int64(n)
	return n, err
}

// ReadAt performs pure independent random-access queries safely
func (f *HTTPFile) ReadAt(p []byte, off int64) (int, error) {

	if len(p) == 0 {
		return 0, nil
	}

	if off >= f.fileSize {
		return 0, io.EOF
	}

	req, err := http.NewRequest("GET", f.URL, nil)

	if err != nil {
		return 0, err
	}

	endOffset := off + int64(len(p)) - 1

	if endOffset >= f.fileSize {
		endOffset = f.fileSize - 1
	}

	rangeHeader := fmt.Sprintf("bytes=%d-%d", off, endOffset)
	req.Header.Set("Range", rangeHeader)
	req.Header.Set("Accept-Encoding", "identity")

	rsp, err := f.Client.Do(req)

	if err != nil {
		return 0, err
	}

	defer rsp.Body.Close()

	if rsp.StatusCode == http.StatusRequestedRangeNotSatisfiable {
		return 0, io.EOF
	}

	if rsp.StatusCode != http.StatusPartialContent && rsp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("remote server error status: %d", rsp.StatusCode)
	}

	bytesToRead := int(endOffset - off + 1)
	n, err := io.ReadFull(rsp.Body, p[:bytesToRead])

	if err == io.ErrUnexpectedEOF && endOffset == f.fileSize-1 {
		return n, nil
	}
	return n, err
}

func (f *HTTPFile) Stat() (fs.FileInfo, error) {
	return f, nil
}

func (f *HTTPFile) Close() error {
	return nil
}

func (f *HTTPFile) Name() string {
	return f.fileName
}

func (f *HTTPFile) Size() int64 {
	return f.fileSize
}

func (f *HTTPFile) Mode() fs.FileMode {
	// Enforce Read-Only permissions flag
	return 0444
}

func (f *HTTPFile) ModTime() time.Time {
	return time.Now()
}

func (f *HTTPFile) IsDir() bool {
	return false
}

func (f *HTTPFile) Sys() any {
	return nil
}
