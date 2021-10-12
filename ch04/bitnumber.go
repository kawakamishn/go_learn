package main

import "fmt"

func main() {
	hoi := [4]byte{1, 2, 3, 4}
	zero(&hoi)
}

func zero(ptr *[4]byte) {
	fmt.Print(*ptr)
	for i := range ptr {
		fmt.Print(ptr[i])
	}
}
