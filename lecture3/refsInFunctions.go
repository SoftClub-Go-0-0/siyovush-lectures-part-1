package main

import "fmt"

func main() {
	var totalScores float64
	var score int
	for score != 9999 {
		fmt.Print("Enter current score: ")
		sumScores(score, &totalScores) // int, 0x982748427
		fmt.Scan(&score)
	}
	fmt.Println("\n\nTotal scores:", totalScores)

	hello("SoftClub")
}

func sumScores(current int, total *float64) {
	*total += float64(current)
}

func hello(name string) {
	fmt.Println("Hello", name)
}
