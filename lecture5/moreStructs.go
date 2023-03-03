package main

import "fmt"

type Car struct {
	ID               uint
	Manufacturer     string
	Model            string
	YearOfProduction int
}

func main() {
	//var car Car

	car := Car{
		ID:               1,
		Manufacturer:     "BMW",
		Model:            "M3",
		YearOfProduction: 1997,
	}

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
