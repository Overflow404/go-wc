package mock

import (
	"go-wc/app"
	"os"
)

type File struct {
	CloseFunc func() error
	ReadFunc  func(p []byte) (int, error)
}

func (c File) Close() error {
	return c.CloseFunc()
}

func (c File) Read(p []byte) (int, error) {
	return c.ReadFunc(p)
}

type FileWrapper struct {
	StatFunc func(string) (os.FileInfo, error)
	OpenFunc func(string) (app.File, error)
}

func (m FileWrapper) Stat(name string) (os.FileInfo, error) {
	return m.StatFunc(name)
}

func (m FileWrapper) Open(name string) (app.File, error) {
	return m.OpenFunc(name)
}
