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

	var isLocMin bool
	var isFirstLocMin = true
	var maxLocMin float64

	for i := 0; i < n; i++ {
		isLocMin = true

		// checking right neighbour
		if i < n-1 {
			if ar[i] > ar[i+1] {
				isLocMin = false
				continue
			}
		}

		// checking left neighbour
		if i > 0 {
			if ar[i] > ar[i-1] {
				isLocMin = false
				continue
			}
		}

		// local minimum values comparison
		if isLocMin {
			if isFirstLocMin {
				maxLocMin = ar[i]
				isFirstLocMin = false
			} else {
				if maxLocMin < ar[i] {
					maxLocMin = ar[i]
				}
			}
		}
	}
	fmt.Println(maxLocMin)
}
