package main

import "os"

type ByteCounter struct{}

func (b ByteCounter) Count(filename string) (int64, error) {
	fileStat, statError := os.Stat(filename)

	if statError != nil {
		return 0, statError
	}

	return fileStat.Size(), nil
}
