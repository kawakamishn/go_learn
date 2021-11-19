// go test -v

package main

import (
	"fmt"
	"math"
)

func main() {
	maximum := max(0, 1, 2, 3.4)
	fmt.Print(maximum)
}

func max(vals ...float64) float64 {
	maximum := math.Inf(-1)
	for _, val := range vals {
		if val > maximum {
			maximum = val
		}
	}
	return maximum
}

func min(vals ...float64) float64 {
	minimum := math.Inf(0)
	for _, val := range vals {
		if val < minimum {
			minimum = val
		}
	}
	return minimum
}
