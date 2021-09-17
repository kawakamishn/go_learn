package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "", counts)
	} else {
		for _, fileName := range files {
			f, err := os.Open(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %/v\n", err)
				continue
			}
			countLines(f, fileName, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, fileName string, counts map[string]int) { // 二番目の引数にfileNameを加えて関数内で辞書のキーに結合させるようにした。
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()+fileName]++
	}
}
