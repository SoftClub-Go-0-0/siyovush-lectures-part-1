package main

import "fmt"

func main() {
	var a, b, c float64
	for i := 0; i < 2; i++ {
		fmt.Print("A = ")
		fmt.Scan(&a)
		fmt.Print("B = ")
		fmt.Scan(&b)
		fmt.Print("C = ")
		fmt.Scan(&c)
		shiftLeft3(&a, &b, &c)
		fmt.Printf("\nA = %.2f\n", a)
		fmt.Printf("B = %.2f\n", b)
		fmt.Printf("C = %.2f\n", c)
	}
}

func shiftLeft3(a, b, c *float64) {
	swap(b, c)
	swap(a, c)
}

func swap(x, y *float64) {
	*x, *y = *y, *x
}