package main

import (
	"testing"
)

func TestByteCounter_Count(t *testing.T) {
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
			name:   "String with Random Content",
			input:  "Hello, world!",
			output: 13,
		},
		{
			name:   "String with Numbers",
			input:  "1234567890",
			output: 10,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ByteCounter{}.Count(test.input)
			if result != test.output {
				t.Errorf("Expected %d bytes, but got %d for input: %s", test.output, result, test.input)
			}
		})
	}
}
