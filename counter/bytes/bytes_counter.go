package bytes

import (
	"go-wc/wrapper"
)

type Counter struct{}

func (b Counter) Count(filename string, fileWrapper wrapper.FileWrapper) (int64, error) {
	fileStat, statError := fileWrapper.Stat(filename)

	if statError != nil {
		return 0, statError
	}

	return fileStat.Size(), nil
}
