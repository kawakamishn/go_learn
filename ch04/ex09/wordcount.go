package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	fileReader, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(fileReader)
	fmt.Print(fileReader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		counts[word]++
	}
	fmt.Print(counts)
	fileReader.Close()
}
