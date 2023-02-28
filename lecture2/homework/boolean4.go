package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Printf("\n(A > 2) and (B <= 3): %t\n", a > 2 && b <= 3)
}