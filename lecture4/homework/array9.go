package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]int, n, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	evenElementsCount := 0

	for i := n - 1; i >= 0; i-- {
		if ar[i]%2 == 0 {
			fmt.Printf("%d ", ar[i])
			evenElementsCount++
		}
	}
	fmt.Printf("\nK = %d\n", evenElementsCount)
}
