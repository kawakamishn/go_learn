// Echo1は、そのコマンドライン因数を表示します。
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
}
