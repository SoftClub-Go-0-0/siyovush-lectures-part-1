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

	step := n + 1
	trigger := -1

	for i := 0; step > 1; i += step * trigger {
		trigger *= -1
		step -= 2
		fmt.Print(ar[i], " ")
		if step != 0 {
			fmt.Print(ar[i+(1*trigger)], " ")
		}
	}
	fmt.Println()
}
