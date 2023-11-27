package wrapper

import "os"

type FileWrapper interface {
	Stat(filename string) (os.FileInfo, error)
	Open(filename string) (File, error)
}

type File interface {
	Close() error
	Read(p []byte) (n int, err error)
}
