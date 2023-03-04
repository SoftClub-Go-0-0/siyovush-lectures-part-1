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

	var isLocMax bool
	var isFirstLocMax = true
	var minLocMax float64

	for i := 0; i < n; i++ {
		isLocMax = true

		// checking right neighbour
		if i < n-1 {
			if ar[i] < ar[i+1] {
				isLocMax = false
				continue
			}
		}

		// checking left neighbour
		if i > 0 {
			if ar[i] < ar[i-1] {
				isLocMax = false
				continue
			}
		}

		// local minimum values comparison
		if isLocMax {
			if isFirstLocMax {
				minLocMax = ar[i]
				isFirstLocMax = false
			} else {
				if minLocMax > ar[i] {
					minLocMax = ar[i]
				}
			}
		}
	}
	fmt.Println(minLocMax)
}
