package counter

import "go-wc/wrapper"

type Counter interface {
	Count(filename string, fileWrapper wrapper.FileWrapper) (int64, error)
}
