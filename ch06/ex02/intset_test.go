package main

import (
	"testing"
)

func TestAddAll(t *testing.T) {
	var s IntSet
	s.AddAll(1, 5, 65)
	if s.String() != "{1 5 65}" {
		t.Error()
	}
}
