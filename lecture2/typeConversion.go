package main

import "fmt"

func main() {
	salary := 2000         // somoni
	salaryIncrease := 24.5 // percents

	fmt.Println("Initial salary:", salary)
	fmt.Println("Salary after increasing:", salary+int(float64(salary)*salaryIncrease/100))
}
