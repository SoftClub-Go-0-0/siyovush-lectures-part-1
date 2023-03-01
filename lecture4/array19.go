package main

import "fmt"

func main() {
	arr := make([]float64, 10)
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}

	var index int
	for i := 1; i < 9; i++ {
		if arr[0] < arr[i] && arr[i] < arr[9] {
			index = i + 1
		}
	}

	fmt.Println(index)
}
