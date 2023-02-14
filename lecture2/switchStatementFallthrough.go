package main

import "fmt"

func main() {
	var grade int
	fmt.Scan(&grade)

	switch grade {
	case 11:
		fmt.Println("You have the knowledge of 11th grade and")
		fallthrough
	case 10:
		fmt.Println("You have the knowledge of 10th grade and")
		fallthrough
	case 9:
		fmt.Println("You have the knowledge of 9th grade and")
		fallthrough
	case 8:
		fmt.Println("You have the knowledge of 8th grade and")
		fallthrough
	case 7:
		fmt.Println("You have the knowledge of 7th grade and")
		fallthrough
	case 6:
		fmt.Println("You have the knowledge of 6th grade and")
		fallthrough
	case 5:
		fmt.Println("You have the knowledge of 5th grade and")
		fallthrough
	case 4:
		fmt.Println("You have the knowledge of 4th grade and")
		fallthrough
	case 3:
		fmt.Println("You have the knowledge of 3rd grade and")
		fallthrough
	case 2:
		fmt.Println("You have the knowledge of 2nd grade and")
		fallthrough
	case 1:
		fmt.Println("You have the knowledge of 1st grade")
	default:
		fmt.Println("error")
	}
}
