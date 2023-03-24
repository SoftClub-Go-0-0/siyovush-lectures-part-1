package main

import (
	"flag"
	"fmt"
)

func main() {
	firstValue := flag.Int("firstValue", 200, "Enter the first value")
	secondValue := flag.Int("secondValue", 200, "Enter the second value")
	operator := flag.String("operator", "+", "Enter the operation")

	flag.Parse()

	switch *operator {
	case "+":
		fmt.Println(*firstValue + *secondValue)
	case "-":
		fmt.Println(*firstValue - *secondValue)
	case "*":
		fmt.Println(*firstValue * *secondValue)
	case "/":
		fmt.Println(*firstValue / *secondValue)
	}
}
