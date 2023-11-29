package app

type ByteCounter struct{}

func (b ByteCounter) Count(filename string, fileWrapper FileWrapper) (int64, error) {
	fileStat, statError := fileWrapper.Stat(filename)

	if statError != nil {
		return 0, statError
	}

	return fileStat.Size(), nil
}
