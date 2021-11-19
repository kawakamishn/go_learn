package main

import (
	"fmt"
	"sort"
)

type PalindromeChecker []rune

func (x PalindromeChecker) Len() int           { return len([]rune(x)) }
func (x PalindromeChecker) Less(i, j int) bool { return x[i] < x[j] }
func (x PalindromeChecker) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) {
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome(PalindromeChecker([]rune("abcdcba"))))
	fmt.Println(IsPalindrome(PalindromeChecker([]rune("abcdecba"))))
}
