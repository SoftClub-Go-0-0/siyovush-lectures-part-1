package main

import "fmt"

func main() {
	var a, b float64
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Printf("\n%d\n", sign(a) + sign(b))
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