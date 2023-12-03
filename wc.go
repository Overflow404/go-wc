package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const bytesFlag = "c"
const linesFlag = "l"
const wordsFlag = "w"
const charactersFlag = "m"

var counters = map[string]Counter{
	bytesFlag:      ByteCounter{},
	linesFlag:      LineCounter{},
	wordsFlag:      WordCounter{},
	charactersFlag: CharCounter{},
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("usage: go-wc <flag> <filename>")
	}

	commandLineArguments := getCommandLineArguments()

	if noCommandLineArgumentsAreProvided(commandLineArguments) {
		handleDefaultCommand(os.Args[1:][0], counters)
	} else {
		handleCustomCommand(os.Args[1:][1], commandLineArguments, counters)
	}
}

func getCommandLineArguments() map[string]*bool {
	flags := map[string]*bool{
		bytesFlag:      flag.Bool(bytesFlag, false, "Flag to enable the byte count"),
		linesFlag:      flag.Bool(linesFlag, false, "Flag to enable the lines count"),
		wordsFlag:      flag.Bool(wordsFlag, false, "Flag to enable the words count"),
		charactersFlag: flag.Bool(charactersFlag, false, "Flag to enable the characters count"),
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

func handleDefaultCommand(fileName string, counters map[string]Counter) {
	result, defaultCommandError := defaultCommand(fileName, counters)

	if defaultCommandError != nil {
		log.Fatalf("%v", defaultCommandError)
	}

	fmt.Println(result)
}

func handleCustomCommand(fileName string, commandLineArguments map[string]*bool, counters map[string]Counter) {
	counterHandler := lookupCounterHandler(commandLineArguments, counters)
	result, customCommandError := customCommand(fileName, counterHandler)

	if customCommandError != nil {
		log.Fatalf("%v", customCommandError)
	}

	fmt.Println(fmt.Sprintf("%d %s", result, fileName))
}

func defaultCommand(filePath string, counters map[string]Counter) (string, error) {
	bytes, bytesError := counters[bytesFlag].Count(filePath)
	if bytesError != nil {
		return "", bytesError
	}

	lines, linesError := counters[linesFlag].Count(filePath)
	if linesError != nil {
		return "", linesError
	}

	words, wordsError := counters[wordsFlag].Count(filePath)
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

	return counters[bytesFlag]
}
