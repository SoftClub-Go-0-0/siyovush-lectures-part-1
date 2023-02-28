package main

import (
	"fmt"
	"math"
)

func main() {
	const pi float64 = 3.14
	var r1, r2 float64
	fmt.Print("R1 = ")
	fmt.Scan(&r1)
	fmt.Print("R2 = ")
	fmt.Scan(&r2)
	var s1, s2, s3 float64
	s1 = pi * math.Pow(r1, 2)
	s2 = pi * math.Pow(r2, 2)
	s3 = s1 - s2
	fmt.Printf("\nS1 = %.2f\n", s1)
	fmt.Printf("S2 = %.2f\n", s2)
	fmt.Printf("S3 = %.2f\n", s3)
}