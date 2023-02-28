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
	var min, max float32
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	if min > c {
		min = c
	}
	if max < c {
		max = c
	}
	fmt.Printf("\nMin: %.2f\n", min)
	fmt.Printf("Max: %.2f\n", max)
}