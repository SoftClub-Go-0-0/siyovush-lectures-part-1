package main

import "fmt"

func main() {
	var mySlice []float64
	var l, r int = 0, 9
	for i := 1; i <= 10; i++ {
		mySlice = append(mySlice, float64(i))
	}
	for i := 0; i < 5; i++ {
		fmt.Print(mySlice[l], " ")
		fmt.Print(mySlice[r], " ")
		l++
		r--
	}
}
