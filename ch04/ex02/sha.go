package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {

	whichSHA := "256"
	input := os.Args[1]
	if len(os.Args) > 2 {
		whichSHA = os.Args[2]
		switch whichSHA {
		case "384":
			hash := sha512.Sum384([]byte(input))
			fmt.Printf("%x", hash)
		case "512":
			hash := sha512.Sum512([]byte(input))
			fmt.Printf("%x", hash)
		}
	}
	hash := sha256.Sum256([]byte(input))
	fmt.Printf("%x", hash)
}
