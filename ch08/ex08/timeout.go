package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func countTime(conn net.Conn, InputCh <-chan bool) {
	ticker := time.NewTicker(time.Second)
	counter := 0
	max := 10
	for {
		select {
		case <-ticker.C:
			counter++
			if counter == max {
				fmt.Println("No input for 10 seconds")
				ticker.Stop()
				conn.Close()
				return
			}
		case <-InputCh: // 入力があればcounterが0に戻る。
			counter = 0
		}
	}
}

func handleConn(c net.Conn) {

	input := bufio.NewScanner(c)
	InputCh := make(chan bool)
	go countTime(c, InputCh)
	for input.Scan() {
		InputCh <- true // 入力があればInputChにtrueが渡される。
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()

}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
