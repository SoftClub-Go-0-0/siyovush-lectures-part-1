package main

import (
	"fmt"
	"time"
)

func main() {
	// Foziljon
	go greeting()

	// Muhammad
	fmt.Println("Another message")
	var s string
	fmt.Scan(&s)
	fmt.Println(s)
}

func greeting() {
	fmt.Println("Hello")
	time.Sleep(time.Second * 2)
	fmt.Println("My name is golang")
	time.Sleep(time.Second * 2)
	fmt.Println("What's your name")
	time.Sleep(time.Second * 2)
}
