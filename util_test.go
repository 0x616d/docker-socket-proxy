package main

import (
	"testing"
)

func TestHasVersionPrefix(t *testing.T) {
	tests := []struct {
		path     string
		expected bool
	}{
		{"/v1.55/volumes", true},
		{"/volumes", false},
		{"/", false},
		{"/v1", true},
	}

	for _, test := range tests {
		result := hasVersionPrefix(test.path)
		if result != test.expected {
			t.Errorf("Expected %v for path %s but got, %v", test.expected, test.path, result)
		}
	}
}

func TestRemoveVersionPrefix(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{"/v1.55/volumes/X/Y/Z", "/volumes/X/Y/Z"},
		{"/volumes/X/Y/Z", "/volumes/X/Y/Z"},
		{"/", "/"},
		{"/v1/", "/"},
	}

	for _, test := range tests {
		result := removeVersionPrefix(test.path)
		if result != test.expected {
			t.Errorf("Expected %s but got %s", test.expected, result)
		}
	}
}
