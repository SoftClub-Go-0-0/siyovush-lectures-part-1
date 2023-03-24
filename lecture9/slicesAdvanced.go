package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s)
	fmt.Println(s[:])
	fmt.Println(s[0:9])
	fmt.Println(s[4:6])
	fmt.Println(s[:5])
	fmt.Println(s[4:])
}
