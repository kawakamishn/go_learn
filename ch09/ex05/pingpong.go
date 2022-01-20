// ./pingpong

package main

import (
	"fmt"
	"time"
)

var ch1 = make(chan string, 1)
var ch2 = make(chan string, 1)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	seconds := 0
	var i int64
	for {
		go func() {

			ch1 <- "Hello"
			for {
				i++
				ch2 <- <-ch1
			}
		}()
		go func() {
			for {
				ch1 <- <-ch2
			}
		}()
		<-ticker.C
		seconds = seconds + 1
		if seconds >= 3 {
			ticker.Stop()
			break
		}
	}

	fmt.Print(i/3, "rounds per second")
}
