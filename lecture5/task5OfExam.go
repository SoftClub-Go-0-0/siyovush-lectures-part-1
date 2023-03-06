package main

import (
	"fmt"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

type Employee struct {
	ID            int
	CompanyID     int
	Name          string
	Surname       string
	Email         string
	DateOfJoining time.Time
	Position      string
	Salary        int
}

func (e *Employee) TotalMoneyEarned() int {
	totalMonths := time.Now().Sub(e.DateOfJoining).Hours() / 24 / 30
	return int(totalMonths) * e.Salary
}

func NewEmployee(companyID int) Employee {
	var employee Employee

	employee.ID = rand.Intn(2000) + 1
	employee.CompanyID = companyID
	employee.Name = RandStringBytes(5)
	employee.Surname = RandStringBytes(15)
	employee.Email = RandStringBytes(10)

	years := rand.Intn(5)
	months := rand.Intn(12)

	employee.DateOfJoining = time.Now().AddDate(-years, -months, 0)
	employee.Position = RandStringBytes(10)
	employee.Salary = rand.Intn(6000) + 1 + 1000

	return employee
}

type Company struct {
	ID                 int
	Name               string
	Location           string
	YearOfFoundation   int
	CurrentYearIncome  int
	PreviousYearIncome int
	Employees          []Employee
}

// TotalMoneySpentOnSalaries возвращает общее количество денег потраченных на оплату труда сотрудников
func (c *Company) TotalMoneySpentOnSalaries() (totalMoney int) {

	return
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
		if employee.Name == name {
			return true
		}
	}
	return false
}

// implement employee

func (c *Company) PrintInfo() {
	fmt.Printf("ID:\t\t\t%d\n", c.ID)
	fmt.Printf("Name:\t\t\t%s\n", c.Name)
	fmt.Printf("Location:\t\t%s\n", c.Location)
	fmt.Printf("Year of foundation:\t%d\n", c.YearOfFoundation)
	fmt.Printf("Current year income:\t%d\n", c.CurrentYearIncome)
	fmt.Printf("Previous year income:\t%d\n", c.PreviousYearIncome)
	fmt.Println("Employees")
	for _, e := range c.Employees {
		fmt.Println(e)
	}
}

func NewCompany() Company {
	var company Company

	company.ID = rand.Intn(60) + 1
	company.Name = RandStringBytes(5)
	company.Location = RandStringBytes(10)
	company.YearOfFoundation = rand.Intn(1000) + 1000
	company.CurrentYearIncome = rand.Intn(1000000) + 1
	company.PreviousYearIncome = rand.Intn(1000000) + 1

	employeesNumber := rand.Intn(50) + 1

	company.Employees = make([]Employee, employeesNumber)
	for i := 0; i < employeesNumber; i++ {
		company.Employees[i] = NewEmployee(company.ID)
	}
	return company
}

func main() {
	company := NewCompany()
	company.PrintInfo()
}
