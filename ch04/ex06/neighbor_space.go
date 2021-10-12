package main

/*
unicodeでスペースとして扱われるもの（'\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).）が並んでいる場合、単一ASCIIスペースに変換する。
*/

import (
	"fmt"
	"unicode"
)

func main() {
	b := []byte("abc\r  \n\rdef\t ghi")
	fmt.Printf("%q\n", string(removeNeighborSpace(b))) // 文字列で表示
	fmt.Printf("%d\n", removeNeighborSpace(b))         // 文字コードで表示
}

func removeNeighborSpace(b []byte) []byte {
	lenB := len(b)
	nRemoved := 0 //除去した数を記録しておく
	for i := 0; i <= lenB-1; i++ {
		j := i - nRemoved
		if unicode.IsSpace(rune(b[j])) {
			if j > 0 && unicode.IsSpace(rune(b[j-1])) { //j番目もj-1番目もスペース判定ならj番目を削除
				b = remove(b, j)
				nRemoved++
			} else {
				b[j] = ' ' //j番目だけがスペース判定ならASCIIのスペースに変換
			}
		} else {
			continue //スペースが絡まなければ何もしない
		}
	}
	return b
}

func remove(slice []byte, k int) []byte { //教科書P104にあるk番目の要素を除去する関数
	copy(slice[k:], slice[k+1:])
	return slice[:len(slice)-1]
}
