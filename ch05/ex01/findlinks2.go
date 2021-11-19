// 実行は　./fetch_ex07 https://golang.org | ./findlinks2
// なぜか実行結果が空になる。なぜなのか。

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println("link", link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" { // 要素ノードでかつ<"a" のとき
		for _, a := range n.Attr { // doc.Attrは[]。n.FirstChildを取って初めて始まる。
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil { // 子ノードがいる場合は子ノードを先に探索する
		visit(links, n.FirstChild)
	}
	if n.NextSibling != nil { // 兄弟ノードがいれば兄弟ノードを探索する
		visit(links, n.NextSibling)
	}
	return links
}
