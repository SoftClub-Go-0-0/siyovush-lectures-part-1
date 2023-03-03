package main

import (
	"fmt"
)

type Manufacturer struct {
	Name     string
	Location string
	CEO      string
}

type Person struct {
	Name    string
	Surname string
	Email   string
}

type Vehicle struct {
	WheelsNumber int
	SeatsNumber  int
}

// Car has Manufacturer
// Car has Person
// Car is Vehicle
type Car struct {
	ID uint
	Vehicle
	Manufacturer     Manufacturer
	Model            string
	YearOfProduction int
	Owner            Person
}

func main() {
	//var car Car
	manufacturer := Manufacturer{
		Name:     "BMW",
		Location: "Germany",
		CEO:      "Mehrdod",
	}

	car := Car{
		ID:               1,
		Manufacturer:     manufacturer,
		Model:            "M3",
		YearOfProduction: 1997,
		Owner: Person{
			Name:    "Siyovush",
			Surname: "Radnoy",
			Email:   "paraparadox@protonmail.com",
		},
		Vehicle: Vehicle{
			WheelsNumber: 4,
			SeatsNumber:  5,
		},
	}

	fmt.Println(car.Model)
	fmt.Println(car.Manufacturer.Name)
	fmt.Println(car.Owner.Name)
	fmt.Println(car.WheelsNumber)
	fmt.Println(car.Vehicle.WheelsNumber)

	/*
		fmt.Println("Using the object")
		fmt.Printf("Type:\t%T\n", car)
		car.PrintCarInfo()

		pointerToCar := &car
		fmt.Println("\nUsing the pointer to an object")
		fmt.Printf("Type:\t%T\n", pointerToCar)
		// These three lines are equivalent:
		pointerToCar.PrintCarInfo() // 1
		//(*pointerToCar).PrintCarInfo() // 2
		//(&car).PrintCarInfo()          // 3

		// Declaring pointers to a struct
		//anotherCar := new(Car)
		//anotherCar := &Car{}

		// copying the value
		secondCar := car
		secondCar.Manufacturer = "Mercedes"
		fmt.Println(car.Manufacturer, secondCar.Manufacturer)

		// copying the address
		theSameCar := &car
		theSameCar.Manufacturer = "Bugatti"
		fmt.Println(car.Manufacturer, theSameCar.Manufacturer)

		secondCar.Manufacturer = "Bugatti"

		// comparing instances of structs
		fmt.Println(car == secondCar, car == *theSameCar)

		fmt.Println(car.Model)
		changeModel(car)
		fmt.Println(car.Model)
		ChangeModel(&car)
		fmt.Println(car.Model)
	*/
}

func (c *Car) PrintCarInfo() {
	fmt.Printf("Car #%d:\n", c.ID)
	fmt.Printf("\tManufacturer:\t\t%s\n", c.Manufacturer)
	fmt.Printf("\tModel:\t\t\t%s\n", c.Model)
	fmt.Printf("\tYear of production:\t%d\n", c.YearOfProduction)
}

func changeModel(car Car) {
	car.Model = "Another model"
}

func ChangeModel(car *Car) {
	car.Model = "Second model"
}
