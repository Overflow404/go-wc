package main

import (
	"testing"
)

func TestLineCounter_Count(t *testing.T) {
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
			name:   "String with One Line",
			input:  "Hello, world!",
			output: 1,
		},
		{
			name:   "String with Multiple Lines",
			input:  "Line 1\nLine 2\nLine 3",
			output: 3,
		},
		{
			name:   "String with Trailing Empty Line",
			input:  "Line 1\nLine 2\nLine 3\n",
			output: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := LineCounter{}.Count(test.input)
			if result != test.output {
				t.Errorf("Expected %d lines, but got %d for input: %s", test.output, result, test.input)
			}
		})
	}
}
