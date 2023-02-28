package main

import "fmt"

func main() {
	var a, b, c, positive_numbers_counter, negative_numbers_counter int
	fmt.Print("A = ")
	fmt.Scan(&a)
	if a > 0 {
		positive_numbers_counter++
	} else if a < 0 {
		negative_numbers_counter++
	}
	
	fmt.Print("B = ")
	fmt.Scan(&b)
	if b > 0 {
		positive_numbers_counter++
	} else if b < 0 {
		negative_numbers_counter++
	}

	fmt.Print("C = ")
	fmt.Scan(&c)
	if c > 0 {
		positive_numbers_counter++
	} else if c < 0 {
		negative_numbers_counter++
	}

	fmt.Printf("\nThe number of positive numbers: %d\n", positive_numbers_counter)
	fmt.Printf("The number of negative numbers: %d\n", negative_numbers_counter)
}