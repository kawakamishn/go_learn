// ./timeout & ./netchat3

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) //ここで標準出力とつなぎっぱなしになる。connがcloseされると下に進む。
		log.Println("done")
		done <- struct{}{} //自身のゴルーチンが終了したことをチャネルに知らせる。
	}()
	mustCopy(conn, os.Stdin) //ここで標準入力とつなぎっぱなしになる
	conn.(*net.TCPConn).CloseWrite()
	<-done // バックグラウンドゴルーチンが終了するのを待っている。
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
