package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() { //引数二つをコマンド入力
	s1 := os.Args[1]
	s2 := os.Args[2]
	countBitDifference(s1, s2)
}

func countBitDifference(s1 string, s2 string) {
	c1 := sha256.Sum256([]byte(s1))
	c2 := sha256.Sum256([]byte(s2))
	if len(c1) != len(c2) { //ハッシュ値の長さが異なることはないと思うが一応確認
		fmt.Print("hashed length is different")
	}
	var count int
	for i, v := range c1 {
		if v != c2[i] {
			count++
		}
	}
	fmt.Printf("%d\n", count)
}
