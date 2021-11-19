// 実行は　./fetch_ex07 https://golang.org | ./findlinks1

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println("link", link)
	}
}

func visit(links []string, n *html.Node) []string {
	z := html.NewTokenizer(os.Stdin)
	tagname, _ := z.TagName()
	fmt.Print("tagname", tagname)
	if n.Type == html.ElementNode && n.Data == "a" { // 要素ノードでかつ<"a" のとき
		for _, a := range n.Attr { // doc.Attrは[]。n.FirstChildを取って初めて始まる。
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling { //nというノードから見た子ノードで始まり、その兄弟ノードを一つずつ、なくなるまで調べるためのfor。
		links = visit(links, c)
	}
	return links
}
