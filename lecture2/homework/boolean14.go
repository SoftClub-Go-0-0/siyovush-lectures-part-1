package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	fmt.Printf("\nExactly one of them is positive: ")
	first_condition := is_positive(a) && !is_positive(b) && !is_positive(c)
	second_condition := !is_positive(a) && is_positive(b) && !is_positive(c)
	third_condition := !is_positive(a) && !is_positive(b) && is_positive(c)
	result := first_condition || second_condition || third_condition
	fmt.Printf("%t\n", result)
}

func is_positive(x int) bool {
	return x > 0
}