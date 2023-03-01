package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	step := n
	trigger := -1

	for i := 0; step != 0; i += step * trigger {
		fmt.Print(ar[i], " ")
		step--
		trigger *= -1
	}
	fmt.Println()
}
