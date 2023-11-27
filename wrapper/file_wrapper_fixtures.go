package wrapper

import (
	"os"
)

type MockedFileWrapper struct {
	StatFunc func(string) (os.FileInfo, error)
	OpenFunc func(string) (File, error)
}

func (m MockedFileWrapper) Stat(name string) (os.FileInfo, error) {
	return m.StatFunc(name)
}

func (m MockedFileWrapper) Open(name string) (File, error) {
	return m.OpenFunc(name)
}

type MockFile struct {
	CloseFunc func() error
	ReadFunc  func(p []byte) (int, error)
}

func (c MockFile) Close() error {
	return c.CloseFunc()
}

func (c MockFile) Read(p []byte) (int, error) {
	return c.ReadFunc(p)
}
