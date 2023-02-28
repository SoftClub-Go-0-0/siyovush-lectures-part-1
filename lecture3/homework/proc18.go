package main

import "fmt"
const pi = 3.14

func main() {
	var r float64
	for i := 0; i < 3; i++ {
		fmt.Print("R = ")
		fmt.Scan(&r)
		fmt.Printf("\n%.2f\n", circleS(r))
	}
}

func circleS(r float64) float64 {
	return pi * r * r
}