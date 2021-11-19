package main

import "testing"

func TestSuperJoin(t *testing.T) {
	if SuperJoin(" ", []string{"a", "b"}) != "a b" {
		t.Error()
	}
	if SuperJoin("-", []string{"a", "b"}, []string{"c", "d"}, []string{"e", "f"}) != "a-b-c-d-e-f" {
		t.Error()
	}
}
