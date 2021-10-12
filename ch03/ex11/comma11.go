package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

var hasDot bool = false

func main() {
	//Args引数が複数入ってきても対応可能
	for i := 1; i < len(os.Args); i++ {
		s := os.Args[i]
		s = parsePlusMinus(s)
		sBeforeDot, sAfterDot := parseDot(s)
		fmt.Printf("%s", comma(sBeforeDot))
		fmt.Printf("%s", ".")
		if hasDot {
			fmt.Printf("%s\n", comma(sAfterDot))
		}
	}
}

func parsePlusMinus(s string) string {
	// プラスかマイナスがあれば表示して符号を除いた文字列を返す
	if s[:1] == "+" || s[:1] == "-" {
		sign := s[:1]
		s = s[1:]
		fmt.Printf("  %s", sign)
	}
	return s
}

func parseDot(s string) (string, string) {
	// ドットの前後で文字列を二つに分けて返す
	if indexOfDot := strings.Index(s, "."); indexOfDot >= 0 {
		hasDot = true
		sBeforeDot := s[:indexOfDot]
		sAfterDot := s[indexOfDot+1:]
		return sBeforeDot, sAfterDot
	}
	return s, s
}

func comma(s string) string {
	var buff bytes.Buffer
	remain := len(s) % 3

	for i, v := range s {

		if !((i-remain)%3 == 0 && i != 0) {
			_, _ = fmt.Fprintf(&buff, "%c", v)
			continue
		}
		buff.WriteByte(',')
		_, _ = fmt.Fprintf(&buff, "%c", v)
	}

	return buff.String()
}
