package main

import (
	"errors"
	"fmt"
)

func getUser(yearOfBirth int, uname string) (name string, age int, isAlive bool, err error) {
	age = 2023 - yearOfBirth
	isAlive = age < 110
	if age < 0 {
		return "", 0, false, errors.New("this person wasn't born yet")
	}

	return uname, age, isAlive, nil
}

func main() {
	var year int
	var name string

	_, err := fmt.Scan(&year, &name)
	if err != nil {
		fmt.Println(err)
		return
	}

	name, age, isAlive, err := getUser(year, name)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(name, age, isAlive)

	//fmt.Println("Scan func result:", n, err)
}
