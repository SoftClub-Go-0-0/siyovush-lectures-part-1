package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, b float32
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	var result float32
	switch n {
	case 1:
		result = a + b
	case 2:
		result = a - b
	case 3:
		result = a * b
	case 4:
		result = a / b
	}
	fmt.Printf("\nResult: %.2f\n", result)
}