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

func Test_LookupCounterHandler_ShouldReturnTheByteCounterByDefault(t *testing.T) {
	args := map[string]*bool{}

	result := lookupCounterHandler(args)

	if _, ok := result.(ByteCounter); !ok {
		t.Errorf("Expected ByteCounter by default, but got %T", result)
	}
}
