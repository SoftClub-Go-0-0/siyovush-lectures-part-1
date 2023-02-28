package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	c := n % 10
	b := n % 100 / 10
	fmt.Printf("\nЕдиницы: %d\n", c)
	fmt.Printf("Дестяки: %d\n", b)
}
