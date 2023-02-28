package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Printf("\nДлина незанятой части: %d\n", a%b)
}
