package main

import (
	"fmt"
)

func main() {
	// declaration using var keyword, type and initial value
	var name string = "Dominic Toretto"
	var age int = 37
	var height, weight float32 = 170.8, 85.6 // declaring multiple variables with the same type
	var isBald bool = true

	fmt.Println("Person:")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height, "cm")
	fmt.Println("Weight:", weight, "kg")
	fmt.Println("Is bald:", isBald)

	fmt.Println("\nFavourite car:")
	// declaration using var keyword and initial value
	var model = "Dodge Charger R/T"
	var productionYear = 1970
	var engineDisplacement = 7.2
	var canFly = false

	fmt.Println("Model:", model)
	fmt.Println("Production year:", productionYear)
	fmt.Println("Engine displacement:", engineDisplacement)
	fmt.Println("Can fly:", canFly)

	fmt.Println("\nRaces:")
	// short declaration via initial value only
	totalRaces := 9248
	racesWon, racesLost := 9247, 1
	// there can be variables of different types in one short declaration
	strongestOpponent, opponentsCar, isCarAmericanProduction := "Brian O'Conner", "Toyota Supra 1994", false

	fmt.Println("Total races:", totalRaces)
	fmt.Println("Races won:", racesWon)
	fmt.Println("Races lost:", racesLost)
	fmt.Println("Strongest opponent:", strongestOpponent)
	fmt.Println("Opponent's car:", opponentsCar)
	fmt.Println("This car is American production:", isCarAmericanProduction)
}
