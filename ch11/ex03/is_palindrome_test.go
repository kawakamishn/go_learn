// 最初の文字と最後の文字は違うはずだが回文がTrueになってしまっている？

package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func randomNonPalindrome(rng *rand.Rand) string {
	n := 7 //回文が偶然発生しやすいように長さを小さくとる
	runes := make([]rune, n)
	for runes[0] == runes[n-1] { //最初と最後が一致していなければOK。でなければやり直し。
		for i := 0; i < n; i++ {
			r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
			runes[i] = r
		}
	}
	if runes[0] == runes[n-1] {
		fmt.Print("yes")
	}
	return string(runes)
}

func TestRandomNonPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = True", p)
		}
	}
}
