package main

import "fmt"

func main() {
	numbers := make(chan int)
	sumOfNumbers := 0

	go func() {
		for {
			sumOfNumbers += <-numbers //
		}
	}()

	for {
		var n int
		fmt.Scan(&n)
		if n == 999 {
			fmt.Println(sumOfNumbers)
			return
		}
		numbers <- n
	}
}
