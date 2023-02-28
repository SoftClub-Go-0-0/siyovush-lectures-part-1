package main

import "fmt"

func main() {
	var y int
	fmt.Print("Year: ")
	fmt.Scan(&y)
	fmt.Printf("\nAge: %d\n", (y - 1) / 100 + 1)
}