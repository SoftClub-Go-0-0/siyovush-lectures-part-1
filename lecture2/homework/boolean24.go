package main

import "fmt"

func main() {
	var a, b, c float32
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	d := b * b - 4 * a * c
	fmt.Printf("\nThe quadratic equeation has real solutions: ")
	fmt.Printf("%t\n", d >= 0)
}