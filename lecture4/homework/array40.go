package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Print("R = ")
	fmt.Scan(&r)

	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	nearest := ar[0]
	minDiff := math.Abs(nearest - r)
	var diff float64

	for i := 1; i < n; i++ {
		diff = math.Abs(ar[i] - r)
		if minDiff > diff {
			minDiff = diff
			nearest = ar[i]
		}
	}

	fmt.Println(nearest)
}
