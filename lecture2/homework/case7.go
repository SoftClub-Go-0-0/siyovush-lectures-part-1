package main

import "fmt"

func main() {
	var weight_unit int
	fmt.Print("Weight unit no.: ")
	fmt.Scan(&weight_unit)
	var weight float32
	fmt.Print("Weight: ")
	fmt.Scan(&weight)
	weight_in_kg := weight
	switch weight_unit {
	case 2:
		weight_in_kg = weight / 1_000_000
	case 3:
		weight_in_kg = weight / 1_000
	case 4:
		weight_in_kg = weight * 1_000
	case 5:
		weight_in_kg = weight * 100
	}
	fmt.Printf("\nWeight in kgs: %f\n", weight_in_kg)
}