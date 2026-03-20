package main

import (
	"fmt"
)

type strl []string

func (s *strl) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *strl) Set(value string) error {
	*s = append(*s, value)
	return nil
}
