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

	var descPartsAmount int

	for i := 0; i < n-1; i++ {
		if ar[i] > ar[i+1] {
			for i < n-1 && ar[i] > ar[i+1] {
				i++
			}
			descPartsAmount++
		}
	}

	fmt.Println(descPartsAmount)
}
