package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	power2 := make([]int, n)
	for i := 0; i < n; i++ {
		power2[i] = int(math.Pow(2, float64(i+1)))
	}
	fmt.Println(power2)
}
