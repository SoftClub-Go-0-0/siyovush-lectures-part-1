package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Hello! What's the current color of the traffic light? (red/yellow/green): ")
	var currentColor string
	fmt.Scan(&currentColor)

	switch currentColor {

	case "red":
		fmt.Println("It's red. Don't pass the road! Wait...")
		for i := 1; i < 4; i++ {
			fmt.Println(i, "...")
			time.Sleep(time.Second)
		}
		fallthrough

	case "yellow":
		fmt.Println("It's yellow. Be ready. Wait...")
		for i := 1; i < 4; i++ {
			fmt.Println(i, "...")
			time.Sleep(time.Second)
		}
		fallthrough

	case "green":
		fmt.Println("It's green. Now go! Hurry...")
		for i := 1; i < 4; i++ {
			fmt.Println(i, "...")
			time.Sleep(time.Second)
		}
		fmt.Println("You're on the other side!")

	default:
		fmt.Println("Dude, is there such a color in a traffic light? :)")
	}
}
