package main

import (
	"fmt"
	"testing"
)

func TestIntersectWith(test *testing.T) {
	var s IntSet
	var t IntSet
	// 同一の長さ（両方128以下）で積集合が取れていること
	s.AddAll(1, 5, 65, 100, 101)
	t.AddAll(1, 65, 100)
	s.IntersectWith(&t)
	if s.String() != "{1 65 100}" {
		test.Error()
	}

	// 異なる長さ（片側だけ65以上を持つ）でも積集合が取れていること
	s.Clear()
	t.Clear()
	s.AddAll(1, 5, 65)
	t.AddAll(1, 5)
	s.IntersectWith(&t)
	if s.String() != "{1 5}" {
		test.Error()
	}
}

func TestDifferencetWith(test *testing.T) {
	var s IntSet
	var t IntSet
	// 同一の長さ（両方128以下）で差集合が取れていること
	s.AddAll(1, 5, 65, 100, 101)
	t.AddAll(1, 65, 100)
	s.DifferenceWith(&t)
	fmt.Print(s.String())
	if s.String() != "{5 101}" {
		test.Error()
	}
}

func TestSymmetricDifference(test *testing.T) {
	var s IntSet
	var t IntSet
	// 同一の長さ（両方128以下）で対称差が取れていること
	s.AddAll(1, 5, 65, 100, 101)
	t.AddAll(1, 65, 100, 103)
	sSym := s.SymmetricDifference(&t)
	fmt.Print(s.String())
	if sSym.String() != "{5 101 103}" {
		test.Error()
	}
}
