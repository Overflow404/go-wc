package main

type Counter interface {
	Count(filename string) int
}
