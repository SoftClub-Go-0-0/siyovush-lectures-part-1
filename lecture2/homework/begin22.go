package main

import "fmt"

func main() {
	var a, b float32
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	/* first way */
	// c := b
	// b = a
	// a = c
	/* second way */
	a, b = b, a
	fmt.Printf("\nA = %.2f\n", a)
	fmt.Printf("B = %.2f\n", b)
}