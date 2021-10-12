package main

import "fmt"

func main() {
	strings := removeDuplication([]string{"abc", "def", "def", "def", "ghi", "def"})
	fmt.Print(strings)
}

func removeDuplication(ss []string) []string {
	lenSS := len(ss)
	nRemoved := 0 //除去した数を記録しておく
	for i := 0; i <= lenSS-2; i++ {
		j := i - nRemoved //除去した分だけStringが短くなるので、指数をずらす。
		if ss[j] == ss[j+1] {
			ss = remove(ss, j+1)
			nRemoved++
		}
	}
	return ss
}

func remove(slice []string, k int) []string { //教科書P104にあるk番目の要素を除去する関数
	copy(slice[k:], slice[k+1:])
	return slice[:len(slice)-1]
}
