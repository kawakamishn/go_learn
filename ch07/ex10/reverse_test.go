package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome(PalindromeChecker([]rune("abcdcba"))) {
		t.Error()
	}
	if !IsPalindrome(PalindromeChecker([]rune("あいうえういあ"))) {
		t.Error()
	}
	if IsPalindrome(PalindromeChecker([]rune("abc"))) {
		t.Error()
	}
}
