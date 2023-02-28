package main

import "fmt"

func main() {
	var a int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Printf("\nA is even: %t\n", a % 2 == 0)
}