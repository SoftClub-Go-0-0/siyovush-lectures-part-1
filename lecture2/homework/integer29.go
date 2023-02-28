package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	cQuantity := (a / c) * (b / c)
	freeSpace := a*b - cQuantity*c*c
	fmt.Printf("\nКоличество квадратов: %d\n", cQuantity)
	fmt.Printf("Площадь неазанятой части: %d\n", freeSpace)
}
