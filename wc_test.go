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
		t.Errorf("Expected ByteCounter, but got %T", result)
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

func Test_LookupCounterHandler_ShouldReturnTheByteCounterByDefault(t *testing.T) {
	args := map[string]*bool{}

	result := lookupCounterHandler(args)

	if _, ok := result.(ByteCounter); !ok {
		t.Errorf("Expected ByteCounter by default, but got %T", result)
	}
}
