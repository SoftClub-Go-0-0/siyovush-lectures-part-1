package main

import "fmt"

func main() {
	var a float64
	fmt.Print("A = ")
	fmt.Scan(&a)
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	product := 1.0
	for i := 0; i < n; i++ {
		product *= a
	}
	fmt.Printf("\nA^N = %.2f\n", product)
}