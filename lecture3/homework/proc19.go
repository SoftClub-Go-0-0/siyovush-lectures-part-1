package main

import "fmt"
const pi = 3.14

func main() {
	var r1, r2 float64
	for i := 0; i < 3; i++ {
		fmt.Print("R1 = ")
		fmt.Scan(&r1)
		fmt.Print("R2 = ")
		fmt.Scan(&r2)
		fmt.Printf("\n%.2f\n", ringS(r1, r2))
	}
}

func circleS(r float64) float64 {
	return pi * r * r
}

func ringS(r1, r2 float64) float64 {
	return circleS(r1) - circleS(r2)
}