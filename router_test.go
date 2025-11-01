package main

import (
	"testing"
)

func TestRouter(t *testing.T) {
	r := newRouter()

	r.insert("GET", "/tasks")
	r.insert("POST", "/services/*")
	r.insert("PUT", "/containers/*/start")
	r.insert("DELETE", "/containers/*/stop")
	r.insert("GET", "/images/*/*/info")
	r.insert("GET", "/images/*/*/info/*/json")
	r.insert("GET", "/system/df")

	tests := []struct {
		method   string
		path     string
		expected bool
	}{
		{"GET", "/tasks", true},
		{"GET", "/tasks/X", false},
		{"POST", "/services/X", true},
		{"POST", "/services/X/Y", false},
		{"PUT", "/containers/X/start", true},
		{"PUT", "/containers/X/start/Y", false},
		{"DELETE", "/containers/X/stop", true},
		{"DELETE", "/containers/X/stop/Y", false},
		{"GET", "/images/X/Y/info", true},
		{"GET", "/images/X/Y/info/X/json", true},
		{"GET", "/images/X/Y/info/X/yaml", false},
		{"GET", "/system/df", true},
		{"POST", "/system/df", false},
		{"PATCH", "/tasks", false},
	}

	for _, test := range tests {
		result := r.search(test.method, test.path)
		if result != test.expected {
			t.Errorf("Method %s, Path %s: expected %v, got %v", test.method, test.path, test.expected, result)
		}
	}
}
