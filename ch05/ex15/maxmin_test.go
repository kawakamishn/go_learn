package main

import "testing"

func TestMax(t *testing.T) {
	if max(-1, 0, 1, 2, 3.4) != 3.4 {
		t.Error(`max(0, 1, 2, 3.4) != 3.4`)
	}
}

func TestMin(t *testing.T) {
	if min(-1, 0, 1, 2, 3.4) != -1 {
		t.Error(`min(-1, 0, 1, 2, 3.4) != -1`)
	}
}
