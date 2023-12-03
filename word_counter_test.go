package main

import (
	"testing"
)

func TestWordCounter_Count(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output int
	}{
		{
			name:   "Empty String",
			input:  "",
			output: 0,
		},
		{
			name:   "String with One Word",
			input:  "Hello!",
			output: 1,
		},
		{
			name:   "String with Multiple Words",
			input:  "This is a test sentence.",
			output: 5,
		},
		{
			name:   "String with Numbers",
			input:  "123 456 789",
			output: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := WordCounter{}.Count(test.input)
			if result != test.output {
				t.Errorf("Expected %d words, but got %d for input: %s", test.output, result, test.input)
			}
		})
	}
}
