package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func main() {
	var s IntSet
	s.Add(3)
	s.Add(1)
	fmt.Print(s.String())
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int { // １が立っている数を調べる
	var length int
	for _, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				length++
			}
		}
	}
	return length
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] -= 1 << bit
}

func (s *IntSet) Clear(x int) {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	var c IntSet
	c.words = make([]uint64, len(s.words))
	copy(c.words, s.words)
	return &c
}
