package main

import "fmt"

func main() {
	direction := 'С'
	fmt.Print("Current direction: ")
	fmt.Scanf("%c", &direction)
	var command int
	fmt.Print("Command: ")
	fmt.Scan(&command)
	switch command {
	case 1:
		switch direction {
		case 'С':
			direction = 'З'
		case 'З':
			direction = 'Ю'
		case 'Ю':
			direction = 'В'
		case 'В':
			direction = 'С'
		}
	case -1:
		switch direction {
		case 'С':
			direction = 'В'
		case 'В':
			direction = 'Ю'
		case 'Ю':
			direction = 'З'
		case 'З':
			direction = 'С'
		}
	}
	fmt.Print(direction)
	fmt.Printf("\nNew direction: %c\n", direction)
}