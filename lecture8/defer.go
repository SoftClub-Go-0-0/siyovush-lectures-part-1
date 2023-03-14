package main

import "fmt"

func main() {
	s := 10
	defer printNumber(s)

	s += 10
	printNumber(s)
}

func printNumber(i int) {
	fmt.Println(i)
}