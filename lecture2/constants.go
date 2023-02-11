package main

// program for calculating the area of a circle:
// s = pi * radius ^ 2

import "fmt"

func main() {
	const pi = 3.14
	var radius float64
	fmt.Print("Радиуси давра (бо метр): r = ")
	fmt.Scan(&radius)
	s := pi * radius * radius
	fmt.Println("Масоҳати ин давра: S =", s, "метри квадратӣ")
}
