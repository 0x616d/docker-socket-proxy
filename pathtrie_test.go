package main

import (
	"testing"
)

func TestPathTrie(t *testing.T) {
	p := newPathTrie()

	tests := []struct {
		path string
	}{
		{"/tasks"},
		{"/services/*"},
		{"/containers/*/start"},
		{"/containers/*/stop"},
		{"/images/*/*/info"},
		{"/images/*/*/info/*/json"},
		{"/system/df"},
	}

	for _, test := range tests {
		p.insert(test.path)
	}

	searchTests := []struct {
		path     string
		expected bool
	}{
		{"/tasks", true},
		{"/tasks/X", false},
		{"/tasks/*", false},
		{"/services/X", true},
		{"/services/X/Y", false},
		{"/containers/X/start", true},
		{"/containers/X/start/Y", false},
		{"/containers/X/stop", true},
		{"/containers/X/stop/Y", false},
		{"/images/X", false},
		{"/images/X/Y", false},
		{"/images/X/Y/Z", false},
		{"/images/X/Y/info", true},
		{"/images/X/Y/info/X/yaml", false},
		{"/images/X/Y/info/X/json", true},
		{"/system", false},
		{"/system/df", true},
		{"/system/df/X", false},
		{"/system/df/X/Y", false},
		{"/system/df/X/Y/Z", false},
		{"/X", false},
		{"/X/Y", false},
		{"/X/Y/Z", false},
	}

	for _, test := range searchTests {
		result := p.search(test.path)
		if result != test.expected {
			t.Errorf("Expected %v for path %s, but got %v", test.expected, test.path, result)
		}
	}
}
