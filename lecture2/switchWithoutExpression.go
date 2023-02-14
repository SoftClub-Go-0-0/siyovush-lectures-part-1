package main

import "fmt"

func main() {
	var average int
	fmt.Scan(&average)

	var absentDays int
	fmt.Scan(&absentDays)

	switch {
	case average < 80:
		//fmt.Println("You have been excluded from the course")
		fallthrough
	case absentDays > 2:
		fmt.Println("You have been excluded from the course")
	default:
		fmt.Println("You can continue the course")
	}
}
