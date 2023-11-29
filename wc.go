package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var CounterMap = map[string]Counter{
	"c": ByteCounter{},
	"l": LineCounter{},
}

func main() {
	if len(os.Args) <= 2 {
		log.Fatalf("usage: go-wc <flag> <filename>")
	}

	filePath := os.Args[1:][1]

	counterHandler := lookupCounterHandler(getCommandLineArguments())

	count, counterError := counterHandler.Count(filePath)

	if counterError != nil {
		log.Fatalf("%v", counterError)
	}

	fmt.Println(fmt.Sprintf("%d %s", count, filePath))
}

func lookupCounterHandler(flags map[string]*bool) Counter {
	for key, value := range flags {
		if *value {
			if counterHandler, ok := CounterMap[key]; ok {
				return counterHandler
			}
		}
	}

	return CounterMap["c"]
}

func getCommandLineArguments() map[string]*bool {
	flags := map[string]*bool{
		"c": flag.Bool("c", false, "Flag to enable the byte count"),
		"l": flag.Bool("l", false, "Flag to enable the lines count"),
	}

	flag.Parse()

	return flags
}
