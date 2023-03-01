package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	for i := 0; i < n; i += 2 {
		fmt.Print(ar[i], " ")
	}
	for i := n - 1 - n%2; i >= 0; i -= 2 {
		fmt.Print(ar[i], " ")
	}
	fmt.Println()
}
