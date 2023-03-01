package main

import "fmt"

func main() {
	var ar [10]int
	for i := 0; i < 10; i++ {
		fmt.Scan(&ar[i])
	}
	a := 0
	for i := 0; i < 9; i++ {
		if ar[i] < ar[9] {
			a = ar[i]
			break
		}
	}
	fmt.Println(a)
}
