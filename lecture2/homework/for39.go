package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Println()
	for i := a; i <= b; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%d\t", i)
		}
	}
	fmt.Println()
}