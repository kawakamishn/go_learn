package main

import (
	"testing"
)

func TestLen(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Add(5)
	s.Add(65)
	if s.Len() != 3 {
		t.Error()
	}
}

func TestRemove(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Add(5)
	s.Add(65)
	s.Remove(5)

	if s.String() != "{1 65}" {
		t.Error()
	}
	s.Remove(65)
	if s.String() != "{1}" {
		t.Error()
	}
}

func TestCopy(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Add(5)
	s.Add(65)
	c := s.Copy()
	if c.String() != "{1 5 65}" {
		t.Error()
	}
}
