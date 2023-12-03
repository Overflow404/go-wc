package main

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

type MockCounter struct {
	CountFunc func(filePath string) (int64, error)
}

func (m MockCounter) Count(filePath string) (int64, error) {
	if m.CountFunc != nil {
		return m.CountFunc(filePath)
	}
	return 0, nil
}

func Test_GetCommandLineArguments_ShouldReturnAMapWithTheBytesFlagSet(t *testing.T) {
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

func Test_NoCommandLineArgumentsAreProvided_ShouldReturnTrueWhenNoArguments(t *testing.T) {
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

	result := noCommandLineArgumentsAreProvided(arguments)

	if result != true {
		t.Errorf("Expected %v got %v", bytesFlag, result)
	}

}

func Test_NoCommandLineArgumentsAreProvided_ShouldReturnFalseWhenAtLeastOneArgument(t *testing.T) {
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

	result := noCommandLineArgumentsAreProvided(arguments)

	if result != false {
		t.Errorf("Expected %v got %v", linesFlag, result)
	}

}

func TestDefaultCommand_EdgeCases(t *testing.T) {
	tests := []struct {
		name          string
		errorCounter  string
		mockedError   error
		expectedError error
	}{
		{
			name:          "ByteCounterError",
			errorCounter:  "c",
			mockedError:   errors.New("error counting bytes"),
			expectedError: errors.New("error counting bytes"),
		},
		{
			name:          "LineCounterError",
			errorCounter:  "l",
			mockedError:   errors.New("error counting lines"),
			expectedError: errors.New("error counting lines"),
		},
		{
			name:          "WordCounterError",
			errorCounter:  "w",
			mockedError:   errors.New("error counting words"),
			expectedError: errors.New("error counting words"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counters := map[string]Counter{
				"c": MockCounter{},
				"l": MockCounter{},
				"w": MockCounter{},
			}

			counters[tt.errorCounter] = MockCounter{
				CountFunc: func(filePath string) (int64, error) {
					return 0, tt.mockedError
				},
			}

			_, defaultCommandError := defaultCommand("mockFilePath", counters)

			if defaultCommandError.Error() != tt.expectedError.Error() {
				t.Errorf("Expected error message '%s', got '%s'", tt.expectedError, defaultCommandError)
			}
		})
	}
}

func TestDefaultCommand_ShouldReturnBytesLinesWordsCount(t *testing.T) {
	counters := map[string]Counter{
		"c": MockCounter{CountFunc: func(filePath string) (int64, error) {
			return 50, nil
		}},
		"l": MockCounter{CountFunc: func(filePath string) (int64, error) {
			return 100, nil
		}},
		"w": MockCounter{CountFunc: func(filePath string) (int64, error) {
			return 200, nil
		}},
	}

	expectedResult := fmt.Sprintf("%d %d %d %s", 100, 200, 50, "mockFilePath")
	result, defaultCommandError := defaultCommand("mockFilePath", counters)

	if defaultCommandError != nil {
		t.Errorf("Unexpected error %v", defaultCommandError)
	}

	if expectedResult != result {
		t.Errorf("Expected %s got %s", expectedResult, result)
	}
}

func TestCustomCommand_ShouldReturnErrorIfCounterFails(t *testing.T) {
	counter := MockCounter{CountFunc: func(filePath string) (int64, error) {
		return 0, errors.New("generic error")
	}}

	expectedCustomCommandError := errors.New("generic error")
	_, customCommandError := customCommand("mockFilePath", counter)

	if customCommandError.Error() != expectedCustomCommandError.Error() {
		t.Errorf("Expected error message %s, got %s", expectedCustomCommandError, customCommandError)
	}
}

func TestCustomCommand_ShouldReturnCountIfSucceed(t *testing.T) {
	counter := MockCounter{CountFunc: func(filePath string) (int64, error) {
		return 500, nil
	}}

	result, customCommandError := customCommand("mockFilePath", counter)

	if customCommandError != nil {
		t.Errorf("Unexpected error %v", customCommandError)
	}

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

func Test_LookupCounterHandler_ShouldDefaultToByteCounterIfNoFlagsProvided(t *testing.T) {
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
