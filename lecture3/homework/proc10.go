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
	swap(&a, &b)
	swap(&c, &d)
	swap(&b, &c)
	fmt.Printf("\nA = %.2f\n", a)
	fmt.Printf("B = %.2f\n", b)
	fmt.Printf("C = %.2f\n", c)
	fmt.Printf("D = %.2f\n", d)
}

func swap(x, y *float64) {
	*x, *y = *y, *x
}