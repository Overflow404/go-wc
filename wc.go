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
	"w": WordCounter{},
	"m": CharCounter{},
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("usage: go-wc <flag> <filename>")
	}

	commandLineArguments := getCommandLineArguments()

	if noCommandLineArgumentsAreProvided(commandLineArguments) {
		filePath := os.Args[1:][0]
		defaultBehaviour(filePath)
	} else {
		filePath := os.Args[1:][1]
		customBehaviour(filePath, commandLineArguments)
	}
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
		"w": flag.Bool("w", false, "Flag to enable the words count"),
		"m": flag.Bool("m", false, "Flag to enable the characters count"),
	}

	flag.Parse()

	return flags
}

func noCommandLineArgumentsAreProvided(flags map[string]*bool) bool {
	for key, value := range flags {
		if *value {
			if _, ok := CounterMap[key]; ok {
				return false
			}
		}
	}

	return true
}

func customBehaviour(filePath string, commandLineArguments map[string]*bool) {
	counterHandler := lookupCounterHandler(commandLineArguments)

	count, counterError := counterHandler.Count(filePath)

	if counterError != nil {
		log.Fatalf("%v", counterError)
	}

	fmt.Println(fmt.Sprintf("%d %s", count, filePath))
}

func defaultBehaviour(filePath string) {
	bytes, bytesError := ByteCounter{}.Count(filePath)
	if bytesError != nil {
		log.Fatalf("%v", bytesError)
	}

	lines, linesError := LineCounter{}.Count(filePath)
	if linesError != nil {
		log.Fatalf("%v", linesError)
	}

	words, wordsError := WordCounter{}.Count(filePath)
	if wordsError != nil {
		log.Fatalf("%v", wordsError)
	}

	fmt.Println(fmt.Sprintf("%d %d %d %s", lines, words, bytes, filePath))
}
