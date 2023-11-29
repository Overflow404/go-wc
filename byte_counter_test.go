package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestByteCounter_Count_StatSucceed(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name          string
		args          args
		expectedBytes int64
		expectedError error
	}{
		{
			"it should count the bytes if the file exist",
			args{
				filename: "assets/test/default.txt",
			},
			342190,
			nil,
		},
		{
			"it should count zero bytes if the file is empty",
			args{
				filename: "assets/test/empty.txt",
			},
			0,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			byteCounter := ByteCounter{}

			actualBytes, countError := byteCounter.Count(tt.args.filename)

			if countError != nil {
				t.Errorf("Count method returned an error: %v", countError)
			}

			if actualBytes != tt.expectedBytes {
				t.Errorf("Expected %d bytes, got %d bytes", tt.expectedBytes, actualBytes)
			}
		})
	}
}

func TestByteCounter_Count_StatFails(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name          string
		args          args
		expectedBytes int64
		expectedError error
	}{
		{
			"it should count zero bytes and return error if the file does not exist",
			args{
				filename: "assets/test/_.txt",
			},
			0,
			errors.New(fmt.Sprintf("stat %s: no such file or directory", "assets/test/_.txt")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			byteCounter := ByteCounter{}

			actualBytes, actualError := byteCounter.Count(tt.args.filename)

			if actualError.Error() != tt.expectedError.Error() {
				t.Fatalf("Expected error {%s}, got {%s}", tt.expectedError, actualError)
			}

			if actualBytes != tt.expectedBytes {
				t.Errorf("Expected %d bytes, got %d bytes", tt.expectedBytes, actualBytes)
			}
		})
	}
}
