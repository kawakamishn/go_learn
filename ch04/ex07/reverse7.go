package main

import "fmt"

func main() {
	b := []byte("abc")
	reverse(b)
	fmt.Printf("%s", b)
}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
