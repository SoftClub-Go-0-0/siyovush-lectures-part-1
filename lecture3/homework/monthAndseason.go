package main

import (
	"fmt"
	"errors"
)

func main() {
	var monthNumber int
	fmt.Scan(&monthNumber)
	month, season, err := monthAndSeason(monthNumber)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(month, season)
}

func monthAndSeason(monthNumber int) (month, season string, err error) {
	switch monthNumber {
	case 1:
		return "January", "Winter", nil
	case 2:
		return "February", "Winter", nil
	case 3:
		return "March", "Spring", nil
	case 4:
		return "April", "Spring", nil
	case 5:
		return "May", "Spring", nil
	case 6:
		return "June", "Summer", nil
	case 7:
		return "July", "Summer", nil
	case 8:
		return "August", "Summer", nil
	case 9:
		return "September", "Autumn", nil
	case 10:
		return "October", "Autumn", nil
	case 11:
		return "November", "Autumn", nil
	case 12:
		return "December", "Winter", nil
	default:
		return "", "", errors.New("unidentified month number")
	}
}