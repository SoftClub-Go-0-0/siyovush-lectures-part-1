package main

import "fmt"

func main() {
	var k int
	for i := 0; i < 5; i++ {
		fmt.Print("K = ")
		fmt.Scan(&k)
		fmt.Printf("\nDigits count: %d\n", digitsCount(k))
	}
}

func digitsCount(k int) int {
	c := 0
	for k > 0 {
		k /= 10
		c++
	}
	return c
}