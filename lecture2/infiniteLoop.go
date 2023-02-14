package main

import "fmt"

func main() {
	count := 0
	for {
		fmt.Println("Enter new student, SKIP to skip or STOP to finish:")
		var input string
		fmt.Scan(&input)
		if input == "STOP" {
			break
		} else if input == "SKIP" {
			continue
		}
		fmt.Println("Adding", input)
		fmt.Println()
		count++
	}
	fmt.Println("Adding students finished")
	fmt.Println("Added", count, "students")
}
