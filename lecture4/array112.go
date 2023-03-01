package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	numbers := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	for j := 1; j < n; j++ {
		for i := 0; i < n-j; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
			}
		}
		fmt.Println(numbers)
	}
}
