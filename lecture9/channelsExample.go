package main

import "fmt"

func main() {
	channel := make(chan string)

	go func() {
		fmt.Println(<-channel)
	}()

	channel <- "Hello"
}
