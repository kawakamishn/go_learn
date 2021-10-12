package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("2 strings are necessary")
		os.Exit(1)
	}

	s1 := os.Args[1]
	s2 := os.Args[2]
	if areAnagrams(s1, s2) {
		fmt.Println("They are anagrams")
	} else {
		fmt.Println("They are not anagrams")
	}
}

func areAnagrams(s1, s2 string) bool {
	if s1 == s2 || len(s1) != len(s2) {
		return false
	}

	s1 = SortString(s1)
	s2 = SortString(s2)
	if s1 == s2 {
		return true
	}
	return false
}

func SortString(w string) string {
	s := strings.Split(w, "") //文字列を1文字ごとの配列にする
	sort.Strings(s)           //並び替える
	return strings.Join(s, "")
}
