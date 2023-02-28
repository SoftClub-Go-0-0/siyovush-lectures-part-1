package main

import "fmt"

func main() {
	var n int
	fmt.Print("N(sec) = ")
	fmt.Scan(&n)
	fmt.Printf("\nN(hrs) = %d\n", n/60/60)
}
