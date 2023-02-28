package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var a, d float64
	fmt.Scan(&a, &d)

	aProgression := make([]float64, n)
	aProgression[0] = a
	for i := 1; i < n; i++ {
		aProgression[i] = aProgression[i-1] + d
	}
	fmt.Println(aProgression)

	for i := 0; i < len(aProgression); i++ {
		fmt.Printf("%.2f\t", aProgression[i])
	}
	fmt.Println()

	//for _, v := range aProgression {
	//	fmt.Printf("%.2f\t", v)
	//}
	//fmt.Println()
}
