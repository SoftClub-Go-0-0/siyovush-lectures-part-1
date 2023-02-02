package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("Enter your name, please: ")
	fmt.Scan(&name)
	fmt.Println("Welcome to the PHP Programming Language Course,", name, "😁")
}
