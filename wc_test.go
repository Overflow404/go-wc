package main

import (
	"fmt"
	"os"
	"testing"
)

type MockCounter struct {
	CountFunc func(filePath string) int
}

func (m MockCounter) Count(filePath string) int {
	return m.CountFunc(filePath)
}

func Test_GetCommandLineArguments_ShouldReturnAMapOnlyWithTheBytesFlagSet(t *testing.T) {
	originalArgs := os.Args

	os.Args = []string{"wc", "-c"}

	defer func() {
		os.Args = originalArgs
	}()

	bytesFlag := true
	linesFlag := false
	wordsFlag := false
	charactersFlag := false

	expectedArguments := map[string]*bool{
		"c": &bytesFlag,
		"l": &linesFlag,
		"w": &wordsFlag,
		"m": &charactersFlag,
	}

	actualArguments := getCommandLineArguments()

	for expectedKey, expectedValue := range expectedArguments {
		actualValue, exists := actualArguments[expectedKey]
		if !exists {
			t.Errorf("Expected argument %s actual arguments", expectedKey)
			continue
		}

		if *expectedValue != *actualValue {
			t.Errorf("Failed for argument %s: expected %v, actual %v", expectedKey, *expectedValue, *actualValue)
		}
	}

}

func Test_NoCommandLineArguments_ShouldReturnTrueWhenNoArguments(t *testing.T) {
	bytesFlag := false
	linesFlag := false
	wordsFlag := false
	charactersFlag := false

	arguments := map[string]*bool{
		"c": &bytesFlag,
		"l": &linesFlag,
		"w": &wordsFlag,
		"m": &charactersFlag,
	}

	result := noCommandLineArguments(arguments)

	if result != true {
		t.Errorf("Expected %v got %v", bytesFlag, result)
	}

}

func Test_NoCommandLineArguments_ShouldReturnFalseWhenAtLeastOneArgument(t *testing.T) {
	bytesFlag := false
	linesFlag := true
	wordsFlag := false
	charactersFlag := false

	arguments := map[string]*bool{
		"c": &bytesFlag,
		"l": &linesFlag,
		"w": &wordsFlag,
		"m": &charactersFlag,
	}

	result := noCommandLineArguments(arguments)

	if result != false {
		t.Errorf("Expected %v got %v", linesFlag, result)
	}

}

func TestDefaultCommand_ShouldReturnBytesLinesWordsCount(t *testing.T) {
	counters := map[string]Counter{
		"c": MockCounter{CountFunc: func(filePath string) int {
			return 50
		}},
		"l": MockCounter{CountFunc: func(filePath string) int {
			return 100
		}},
		"w": MockCounter{CountFunc: func(filePath string) int {
			return 200
		}},
	}

	expectedResult := fmt.Sprintf("%d %d %d %s", 100, 200, 50, "mockFilePath")
	result := defaultCommand("mockFilePath", "mockedInput", counters)

	if expectedResult != result {
		t.Errorf("Expected %s got %s", expectedResult, result)
	}
}

func TestCustomCommand_ShouldReturnCount(t *testing.T) {
	counter := MockCounter{CountFunc: func(filePath string) int {
		return 500
	}}

	result := customCommand("mockFilePath", counter)

	if result != 500 {
		t.Errorf("Expected %d got %d", 500, result)
	}
}

func Test_LookupCounterHandler_ShouldReturnTheProperCounterHandler(t *testing.T) {
	bytesFlag := false
	linesFlag := true
	wordsFlag := false
	charactersFlag := false

	flags := map[string]*bool{
		"c": &bytesFlag,
		"l": &linesFlag,
		"w": &wordsFlag,
		"m": &charactersFlag,
	}

	counters := map[string]Counter{
		"c": ByteCounter{},
		"l": LineCounter{},
		"w": WordCounter{},
		"m": CharCounter{},
	}

	result := lookupCounterHandler(flags, counters)

	if _, ok := result.(LineCounter); !ok {
		t.Errorf("Expected LineCounter, but got %T", result)
	}
}

func Test_LookupCounterHandler_ShouldDefaultToByteCounterHandlerIfNoFlagsAreProvided(t *testing.T) {
	bytesFlag := false
	linesFlag := false
	wordsFlag := false
	charactersFlag := false

	flags := map[string]*bool{
		"c": &bytesFlag,
		"l": &linesFlag,
		"w": &wordsFlag,
		"m": &charactersFlag,
	}

	counters := map[string]Counter{
		"c": ByteCounter{},
		"l": LineCounter{},
		"w": WordCounter{},
		"m": CharCounter{},
	}

	result := lookupCounterHandler(flags, counters)

	if _, ok := result.(ByteCounter); !ok {
		t.Errorf("Expected ByteCounter, but got %T", result)
	}
}
