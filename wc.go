package main

import (
	"flag"
	"fmt"
	"io"
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
	var input string
	var fileName string

	if inputIsComingFromPipe() {
		input = readPipeInput(input)
	} else {
		if moreArgumentsAreProvided() {
			input = readFileInput(readFileName())
		} else {
			log.Fatalf("usage: go-wc <flag> <filename>")
		}
	}

	commandLineArguments := getCommandLineArguments()

	if noCommandLineArguments(commandLineArguments) {
		processDefaultCommand(fileName, input, counters)
	} else {
		processCustomCommand(fileName, input, commandLineArguments, counters)
	}
}

func readFileName() string {
	return os.Args[len(os.Args)-1]
}

func inputIsComingFromPipe() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func readPipeInput(input string) string {
	pipeInput, readAllError := io.ReadAll(os.Stdin)

	if readAllError != nil {
		log.Fatalf("Error reading through pipe: %v", readAllError)
	}

	input = string(pipeInput)
	return input
}

func moreArgumentsAreProvided() bool {
	return len(os.Args) > 1
}

func readFileInput(fileName string) string {
	fileInput, readFileError := os.ReadFile(fileName)

	if readFileError != nil {
		log.Fatalf("Error reading from file: %v", readFileError)
	}

	return string(fileInput)
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

func noCommandLineArguments(commandLineArguments map[string]*bool) bool {
	for _, value := range commandLineArguments {
		if *value == true {
			return false
		}
	}

	return true
}

func processDefaultCommand(fileName string, input string, counters map[string]Counter) {
	result := defaultCommand(fileName, input, counters)
	fmt.Println(fmt.Sprintf("\t%s", result))
}

func defaultCommand(filePath string, input string, counters map[string]Counter) string {
	bytes := counters[bytesFlag].Count(input)

	lines := counters[linesFlag].Count(input)

	words := counters[wordsFlag].Count(input)

	return fmt.Sprintf("%d %d %d %s", lines, words, bytes, filePath)
}

func processCustomCommand(fileName string, input string, commandLineArguments map[string]*bool, counters map[string]Counter) {
	counterHandler := lookupCounterHandler(commandLineArguments, counters)
	result := customCommand(input, counterHandler)

	fmt.Println(fmt.Sprintf("\t%d %s", result, fileName))
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

func customCommand(input string, counter Counter) int {
	return counter.Count(input)
}
