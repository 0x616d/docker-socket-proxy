package main

import (
	"strings"
)

func hasVersionPrefix(s string) bool {
	return len(s) > 2 && s[1] == 'v' && s[2] == '1'
}

func removeVersionPrefix(path string) string {
	if  hasVersionPrefix(path) {
		return "/" + strings.Join(strings.Split(path, "/")[2:], "/")
	}
	return path
}
