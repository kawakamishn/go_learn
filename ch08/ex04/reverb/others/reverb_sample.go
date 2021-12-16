package main

import (
	"fmt"
	"time"
)

func main() {
	go hello() // hello()関数を別処理として実行
	go goodbye()
	fmt.Println("main1")
	time.Sleep(time.Millisecond * 30)
	fmt.Println("main2")
}

func hello() {
	time.Sleep(time.Millisecond * 10)
	fmt.Println("hello1")
	time.Sleep(time.Millisecond * 10) //追加
	fmt.Println("hello2")             //追加
}

func goodbye() {
	time.Sleep(time.Millisecond * 10)
	fmt.Println("goodbye1")
	time.Sleep(time.Millisecond * 10) //追加
	fmt.Println("goodbye2")           //追加
}
