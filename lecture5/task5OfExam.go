package main

import (
	"fmt"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

type Company struct {
	Name               string
	Location           string
	YearOfFoundation   int
	CurrentYearIncome  int
	PreviousYearIncome int
	Employees          []string
}

// YearsOfExistence возвращает разницу между текущим годом и годом основания компании
func (c *Company) YearsOfExistence() int {
	return 2023 - c.YearOfFoundation
}

// IncomeDifference возвращает разницу между доходом текущего и предыдущего года
func (c *Company) IncomeDifference() int {
	return c.CurrentYearIncome - c.PreviousYearIncome
}

// HasEmployee принимает имя сотрудника и проверяет есть ли такой сотрудник в массиве сотрудников компании
func (c *Company) HasEmployee(name string) bool {
	for _, employee := range c.Employees {
		if employee == name {
			return true
		}
	}
	return false
}

// implement employee

func (c *Company) PrintInfo() {
	fmt.Printf("Name:\t\t\t%s\n", c.Name)
	fmt.Printf("Location:\t\t%s\n", c.Location)
	fmt.Printf("Year of foundation:\t%d\n", c.YearOfFoundation)
	fmt.Printf("Current year income:\t%d\n", c.CurrentYearIncome)
	fmt.Printf("Previous year income:\t%d\n", c.PreviousYearIncome)
}

func NewCompany() Company {
	var company Company

	company.Name = RandStringBytes(5)
	company.Location = RandStringBytes(10)
	company.YearOfFoundation = rand.Intn(1000) + 1000
	company.CurrentYearIncome = rand.Intn(1000000) + 1
	company.PreviousYearIncome = rand.Intn(1000000) + 1

	employeesNumber := rand.Intn(50) + 1

	company.Employees = make([]string, employeesNumber)
	for i := 0; i < employeesNumber; i++ {
		company.Employees[i] = RandStringBytes(15)
	}
	return company
}

func main() {
	company := NewCompany()
	company.PrintInfo()
}
