package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go echo(c, input.Text(), 3*time.Second, &wg)
	}
	go func() {
		wg.Wait()
		c.(*net.TCPConn).CloseWrite()
	}()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
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
