package main

import (
	"fmt"
	"strings"
	"time"
)

func echo(shout string, delay time.Duration) {
	fmt.Println(strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Println(shout)
	time.Sleep(delay)
	fmt.Println(strings.ToLower(shout))
}

func handleConn() {
	shouts := []string{"a", "b"}
	for _, shout := range shouts {
		go echo(shout, 1*time.Second)
	}
}

func main() {
	handleConn()
	time.Sleep(5 * time.Second)
}
