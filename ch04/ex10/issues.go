// buildするときは go run issues.go github.go
// 実行は例えば　./issues lions
// Titleだけ表示している。

package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range result.Items {
		daysPassed := time.Since(item.CreatedAt).Hours() / 24
		if daysPassed < 31 {
			fmt.Print("within one month: ")
			fmt.Print(item.Title)
		} else if daysPassed < 365 {
			fmt.Print("from one month to one year: ")
			fmt.Print(item.Title)
		} else {
			fmt.Print("over one year: ")
			fmt.Print(item.Title)
		}
		fmt.Printf("%s", "\n")
	}
}
