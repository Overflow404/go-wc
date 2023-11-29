package main

import (
	"flag"
	"go-wc/app"
	"os"
	"testing"
)

var applicationName = "word_count.go"
var filePath = "assets/test/default.txt"

func TestGetCommandLineArguments_ShouldSetTheBytesCountPointer(t *testing.T) {
	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	}()

	os.Args = []string{applicationName, "-c", filePath}

	args := getCommandLineArguments()

	if !*args.flagEnabled["c"] {
		t.Errorf("Expected bytesCountFlag to be true, but it was false")
	}
}

func TestGetCommandLineArguments_ShouldSetTheLinesCountPointer(t *testing.T) {
	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	}()

	os.Args = []string{applicationName, "-l", filePath}

	args := getCommandLineArguments()

	if !*args.flagEnabled["l"] {
		t.Errorf("Expected linesCountFlag to be true, but it was false")
	}
}

func TestLookupCounterHandler_ShouldReturnBytesHandler(t *testing.T) {
	args := commandLineArguments{
		flagEnabled: map[string]*bool{
			"c": boolPointer(true),
			"l": boolPointer(true),
		},
	}

	counterHandler := lookupCounterHandler(args)

	_, isBytesCounter := counterHandler.(app.ByteCounter)
	if !isBytesCounter {
		t.Errorf("Expected BytesCounter, but got %T", counterHandler)
	}
}

func TestLookupCounterHandler_ShouldReturnLinesHandler(t *testing.T) {
	args := commandLineArguments{
		flagEnabled: map[string]*bool{
			"c": boolPointer(false),
			"l": boolPointer(true),
		},
	}

	counterHandler := lookupCounterHandler(args)

	_, isLinesCounter := counterHandler.(app.LineCounter)
	if !isLinesCounter {
		t.Errorf("Expected LinesCounter, but got %T", counterHandler)
	}
}

func TestLookupCounterHandler_ShouldReturnDefaultHandler(t *testing.T) {
	args := commandLineArguments{
		flagEnabled: map[string]*bool{
			"c": boolPointer(false),
			"l": boolPointer(false),
		},
	}

	counterHandler := lookupCounterHandler(args)

	_, isBytesCounter := counterHandler.(app.ByteCounter)
	if !isBytesCounter {
		t.Errorf("Expected BytesCounter, but got %T", counterHandler)
	}
}

func boolPointer(b bool) *bool {
	return &b
}
