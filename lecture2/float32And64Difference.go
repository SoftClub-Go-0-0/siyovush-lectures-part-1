package main

import "fmt"

func main() {
	fmt.Println("Difference between float32 and float64")

	var first float32 = 12.123456_789012_345678_901234_567890_123456_789012_345678
	var second float64 = 12.123456_789012_345678_901234_567890_123456_789012_345678

	fmt.Printf("First variable with the precision of 50 digits: %.50f\n", first)
	fmt.Printf("Second variable with the precision of 50 digits: %.50f\n", second)
}
