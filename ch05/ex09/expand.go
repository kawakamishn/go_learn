package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(expand("abcde$fooghi", To123))
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}

func To123(s string) string {
	return "123"
}
