package main

import "fmt"

func area(r float64) float64 {
	return r * r * 3.14
}

func main() {
	var r1, r2, r3 float64
	fmt.Scan(&r1, &r2, &r3)
	fmt.Println(area(r1), area(r2), area(r3))
}
