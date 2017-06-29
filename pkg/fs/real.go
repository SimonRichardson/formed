package fs

import (
	"io"
	"os"
)

type realFilesystem struct{}

// New yields a real disk filesystem
func New() Filesystem {
	return realFilesystem{}
}

// Create takes a path, creates the file and then returns a File back that
// can be used. This returns an error if the file can not be created in
// some way.
func (realFilesystem) Create(path string) (file File, err error) {
	var f *os.File
	f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return
	}

	return realFile{
		File:   f,
		Reader: f,
		Closer: f,
	}, nil
}

// Open takes a path, opens a potential file and then returns a File if
// that file exists, otherwise it returns an error if the file wasn't found.
func (realFilesystem) Open(path string) (file File, err error) {
	var f *os.File
	f, err = os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return
	}

	return realFile{
		File:   f,
		Reader: f,
		Closer: f,
	}, nil
}

// Exists takes a path and checks to see if the potential file exists or
// not.
// Note: If there is an error trying to read that file, it will return false
// even if the file already exists.
func (realFilesystem) Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

type realFile struct {
	*os.File
	io.Reader
	io.Closer
}

// Read reads a file and places the values onto the byte slice. It returns the
// amount that is read (int), but also returns an error if something went wrong
// whilst reading the file.
func (f realFile) Read(p []byte) (int, error) {
	return f.Reader.Read(p)
}

// Close closes the file once done, otherwise it returns an error.
func (f realFile) Close() error {
	return f.Closer.Close()
}
