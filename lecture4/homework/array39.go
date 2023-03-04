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

	var monotonicPartsAmount int

	for i := 0; i < n-1; {
		// Descending order
		if ar[i] > ar[i+1] {
			for i < n-1 && ar[i] > ar[i+1] {
				i++
			}
			monotonicPartsAmount++
		}

		// Ascending order
		if i < n-1 && ar[i] < ar[i+1] {
			for i < n-1 && ar[i] < ar[i+1] {
				i++
			}
			monotonicPartsAmount++
		}

		// If they are equal, i will not change,
		// so we do it ourselves
		if i < n-1 && ar[i] == ar[i+1] {
			i++
		}
	}

	fmt.Println(monotonicPartsAmount)
}
