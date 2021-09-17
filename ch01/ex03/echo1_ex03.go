// Echo1は、そのコマンドライン因数を表示します。
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
}
