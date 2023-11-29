package main

import (
	"flag"
	"fmt"
	"go-wc/app"
	"log"
	"os"
)

type commandLineArguments struct {
	flagEnabled map[string]*bool
}

var CounterMap = map[string]app.Counter{
	"c": app.ByteCounter{},
	"l": app.LineCounter{},
}

type OsFileWrapper struct{}

func (m OsFileWrapper) Stat(filename string) (os.FileInfo, error) {
	return os.Stat(filename)
}

func (m OsFileWrapper) Open(filename string) (app.File, error) {
	return os.Open(filename)
}

func main() {
	if len(os.Args) <= 2 {
		log.Fatalf("usage: go-wc <flag> <filename>")
	}

	filePath := os.Args[1:][1]

	counterHandler := lookupCounterHandler(getCommandLineArguments())

	count, counterError := counterHandler.Count(filePath, OsFileWrapper{})

	if counterError != nil {
		log.Fatalf("%v", counterError)
	}

	fmt.Println(fmt.Sprintf("%d %s", count, filePath))
}

func lookupCounterHandler(arguments commandLineArguments) app.Counter {
	for key, value := range arguments.flagEnabled {
		if *value {
			if counterHandler, ok := CounterMap[key]; ok {
				return counterHandler
			}
		}
	}

	return CounterMap["c"]
}

func getCommandLineArguments() commandLineArguments {
	flags := map[string]*bool{
		"c": flag.Bool("c", false, "Flag to enable the byte count"),
		"l": flag.Bool("l", false, "Flag to enable the lines count"),
	}

	flag.Parse()

	return commandLineArguments{
		flagEnabled: flags,
	}
}
