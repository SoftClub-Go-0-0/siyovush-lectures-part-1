package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	oddNumbers := make([]int, n)
	for i := 0; i < n; i++ {
		oddNumbers[i] = i*2 + 1
	}
	fmt.Println(oddNumbers)

	//fmt.Printf("%v\t%T\n", oddNumbers, oddNumbers)
	//fmt.Printf("%v\t%v\n", len(oddNumbers), cap(oddNumbers))
	//
	//oddNumbers = append(oddNumbers, 999, 123, 777)
	//
	//fmt.Printf("%v\t%T\n", oddNumbers, oddNumbers)
	//fmt.Printf("%v\t%v\n", len(oddNumbers), cap(oddNumbers))
	//fmt.Printf("%v\t%v\n", oddNumbers[8], oddNumbers[9])

	//array := [10]int{}
	//array2 := [9]int{}
	//fmt.Printf("%v\t%T\n", array, array)
	//fmt.Printf("%v\t%T\n", array2, array2)
}
