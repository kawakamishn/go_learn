package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6} // 配列のため固定長
	reverse(&a)
	fmt.Println(a)
}

func reverse(s *[6]int) { //入力は配列のポインタ
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		fmt.Print(s[i])
		s[i], s[j] = s[j], s[i]
	}
}
