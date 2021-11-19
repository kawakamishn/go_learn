// go test -v

package main

import (
	"fmt"
	"strings"
)

func main() {
	//con := SuperJoin("-", []string{"a", "b"}, []string{"c", "d"})
	fmt.Print(11 & 12)
	// fmt.Print(con)
}

func SuperJoin(sep string, strs ...[]string) string {
	superJoined := ""
	for _, str := range strs {
		joined := strings.Join(str, sep)
		if superJoined == "" {
			superJoined = joined
			continue
		}
		superJoined = strings.Join([]string{superJoined, joined}, sep)
	}
	return superJoined
}
