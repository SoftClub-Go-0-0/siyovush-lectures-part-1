package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	a := n / 1000
	b := n / 100 % 10
	c := n % 100 / 10
	d := n % 10
	fmt.Printf("\nThe N is palindrome: ")
	fmt.Printf("%t\n", a == d && b == c)
}