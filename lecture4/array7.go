package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	numbers := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%.2f\t", numbers[i])
	}
	fmt.Println()
}
