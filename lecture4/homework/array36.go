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

	var isLocMax, isLocMin bool
	var isFirstMax = true
	var maxNotLocMinOrLocMax float64

	for i := 0; i < n; i++ {
		isLocMax = true
		isLocMin = true

		// checking right neighbour
		if i < n-1 {
			if ar[i] < ar[i+1] {
				isLocMax = false
			}
			if ar[i] > ar[i+1] {
				isLocMin = false
			}
		}

		// checking left neighbour
		if i > 0 {
			if ar[i] < ar[i-1] {
				isLocMax = false
			}
			if ar[i] > ar[i-1] {
				isLocMin = false
			}
		}

		// if it's neither local minimum nor local maximum
		// then compare its value with other ones
		if !isLocMin && !isLocMax {
			if isFirstMax {
				maxNotLocMinOrLocMax = ar[i]
				isFirstMax = false
			} else {
				if maxNotLocMinOrLocMax < ar[i] {
					maxNotLocMinOrLocMax = ar[i]
				}
			}
		}
	}
	fmt.Println(maxNotLocMinOrLocMax)
}
