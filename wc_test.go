package main

import (
	"testing"
)

func Test_LookupCounterHandler_ShouldReturnTheByteCounterHandler(t *testing.T) {
	args := map[string]*bool{
		"c": new(bool),
	}

	*args["c"] = true

	result := lookupCounterHandler(args)

	if _, ok := result.(ByteCounter); !ok {
		t.Errorf("Expected ByteCounter, but got %T", result)
	}
}

func Test_LookupCounterHandler_ShouldReturnTheLineCounterHandler(t *testing.T) {
	args := map[string]*bool{
		"l": new(bool),
	}

	*args["l"] = true

	result := lookupCounterHandler(args)

	if _, ok := result.(LineCounter); !ok {
		t.Errorf("Expected LineCounter, but got %T", result)
	}
}

func Test_LookupCounterHandler_ShouldReturnTheWordCounterHandler(t *testing.T) {
	args := map[string]*bool{
		"w": new(bool),
	}

	*args["w"] = true

	result := lookupCounterHandler(args)

	if _, ok := result.(WordCounter); !ok {
		t.Errorf("Expected WordCounter, but got %T", result)
	}
}

func Test_LookupCounterHandler_ShouldReturnTheCharCounterHandler(t *testing.T) {
	args := map[string]*bool{
		"m": new(bool),
	}

	*args["m"] = true

	result := lookupCounterHandler(args)

	if _, ok := result.(CharCounter); !ok {
		t.Errorf("Expected CharCounter, but got %T", result)
	}
}

func Test_NoCommandLineArgumentsAreProvided_ShouldReturnFalse(t *testing.T) {
	args := map[string]*bool{
		"m": new(bool),
	}

	*args["m"] = true

	result := noCommandLineArgumentsAreProvided(args)

	if result != false {
		t.Errorf("Expected false, but got %T", result)
	}
}

func Test_NoCommandLineArgumentsAreProvided_ShouldReturnTrue(t *testing.T) {
	args := map[string]*bool{}

	result := noCommandLineArgumentsAreProvided(args)

	if result != true {
		t.Errorf("Expected true, but got %T", result)
	}
}
