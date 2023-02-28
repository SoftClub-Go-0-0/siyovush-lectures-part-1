package main

import "fmt"

func main() {
	var k int
	for i := 0; i < 5; i++ {
		fmt.Print("K = ")
		fmt.Scan(&k)
		for n := 0; n < 5; n++ {
			fmt.Printf("%d\t", digitN(k, n))
		}
		fmt.Println()
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

func digitN(k, n int) int {
	digits_count := digitsCount(k)
	if digits_count <= n {
		return -1
	}
	var digit_n int
	for i := 0; i < digits_count; i++ {
		if i == n {
			digit_n = k % 10
		}
		k /= 10
	}
	return digit_n
}