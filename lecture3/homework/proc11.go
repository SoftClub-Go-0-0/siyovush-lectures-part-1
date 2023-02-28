package main

import "fmt"

func main() {
	var a, b, c, d float64
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	fmt.Print("D = ")
	fmt.Scan(&d)
	minmax(&a, &b)
	minmax(&c, &d)
	minmax(&b, &d)
	minmax(&a, &c)
	fmt.Printf("\nMin = %.2f\n", a)
	fmt.Printf("Max = %.2f\n", d)
}

func minmax(x, y *float64) {
	if *x > *y {
		*x, *y = *y, *x
	}
}