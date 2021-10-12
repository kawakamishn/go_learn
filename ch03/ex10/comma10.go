package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	//Args引数が複数入ってきても対応可能
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buff bytes.Buffer
	remain := len(s) % 3 //右から順に3ごとに区切りたいので３で割った端数を除くための値

	for i, v := range s { // バイトごとのループなので、あいうえお → あ,い,う,え,お
		if !((i-remain)%3 == 0 && i != 0) {
			_, _ = fmt.Fprintf(&buff, "%c", v)
			continue
		}
		buff.WriteByte(',')
		_, _ = fmt.Fprintf(&buff, "%c", v)
	}

	return buff.String()
}
