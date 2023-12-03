package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("usage: go-wc <flag> <filename>")
	}

	commandLineArguments := getCommandLineArguments()

	counters := map[string]Counter{
		"c": ByteCounter{},
		"l": LineCounter{},
		"w": WordCounter{},
		"m": CharCounter{},
	}

	if noCommandLineArgumentsAreProvided(commandLineArguments) {
		result, defaultCommandError := defaultCommand(os.Args[1:][0], counters)

		if defaultCommandError != nil {
			log.Fatalf("%v", defaultCommandError)
		}

		fmt.Println(result)
	} else {
		fileName := os.Args[1:][1]
		counterHandler := lookupCounterHandler(commandLineArguments, counters)
		result, customCommandError := customCommand(fileName, counterHandler)

		if customCommandError != nil {
			log.Fatalf("%v", customCommandError)
		}

		fmt.Println(fmt.Sprintf("%d %s", result, fileName))
	}
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
	for _, value := range flags {
		if *value {
			return false
		}
	}

	return true
}

func defaultCommand(filePath string, counters map[string]Counter) (string, error) {
	bytes, bytesError := counters["c"].Count(filePath)
	if bytesError != nil {
		return "", bytesError
	}

	lines, linesError := counters["l"].Count(filePath)
	if linesError != nil {
		return "", linesError
	}

	words, wordsError := counters["w"].Count(filePath)
	if wordsError != nil {
		return "", wordsError
	}

	return fmt.Sprintf("%d %d %d %s", lines, words, bytes, filePath), nil
}

func customCommand(filePath string, counter Counter) (int64, error) {
	count, counterError := counter.Count(filePath)

	if counterError != nil {
		return 0, counterError
	}

	return count, nil
}

func lookupCounterHandler(flags map[string]*bool, counters map[string]Counter) Counter {
	for key, value := range flags {
		if *value {
			if counterHandler, ok := counters[key]; ok {
				return counterHandler
			}
		}
	}

	return counters["c"]
}
