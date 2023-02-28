package main

import "fmt"

func main() {
	var a, b, c float64
	for i := 0; i < 3; i++ {
		fmt.Print("A = ")
		fmt.Scan(&a)
		fmt.Print("B = ")
		fmt.Scan(&b)
		fmt.Print("C = ")
		fmt.Scan(&c)
		fmt.Printf("\n%d\n", rootCount(a, b, c))
	}
}

func rootCount(a, b, c float64) int {
	d := b * b - 4 * a * c
	return sign(d) + 1
}

func sign(x float64) int {
	if x < 0 {
		return -1
	} else if x == 0 {
		return 0
	} else {
		return 1
	}
}