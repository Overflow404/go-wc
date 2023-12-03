package main

type ByteCounter struct{}

func (b ByteCounter) Count(input string) int {
	return len([]byte(input))
}
