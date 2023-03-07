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

func (c Company) String() string {
	return fmt.Sprintf(
		"\tID: %d\n\tName: %s\n\tLocation: %s\n\tYear of foundation: %d\n\tNumber of employees: %d\n",
		c.ID, c.Name, c.Location, c.YearOfFoundation, len(c.Employees),
	)
}

func main() {
	company := NewCompany()
	fmt.Println("Company")
	fmt.Print(company)
}
