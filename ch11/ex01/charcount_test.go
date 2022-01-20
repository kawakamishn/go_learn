package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCharcount(t *testing.T) {
	tests := []struct {
		input   string
		counts  map[rune]int
		utflen  map[int]int
		invalid int
	}{
		{
			input:   "ABC あいう,.",
			counts:  map[rune]int{'A': 1, 'B': 1, 'C': 1, ' ': 1, 'あ': 1, 'い': 1, 'う': 1, ',': 1, '.': 1},
			utflen:  map[int]int{1: 6, 3: 3},
			invalid: 0,
		},
	}
	for _, test := range tests {
		counts, utflen, invalid := charCount(strings.NewReader(test.input))
		if !reflect.DeepEqual(counts, test.counts) {
			t.Errorf("%q runes: got %v, want %v", test.input, counts, test.counts)
		}
		if !reflect.DeepEqual(utflen, test.utflen) {
			t.Errorf("%q uftlen: got %v, want %v", test.input, utflen, test.utflen)
		}
		if !reflect.DeepEqual(invalid, test.invalid) {
			t.Errorf("%q invalid: got %v, want %v", test.input, invalid, test.invalid)
		}
	}
}
