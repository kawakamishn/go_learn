package main

import (
	"io"
	"log"
	"net"
	"time"
)

const (
	portnumber = "8000"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second) //これがないと1秒間隔で表示してくれない
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:"+portnumber)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept() // この部分で待つ
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
