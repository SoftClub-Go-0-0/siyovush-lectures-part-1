package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	orderCh := make(chan bool)
	applicationCh := make(chan bool)
	reviewCh := make(chan bool)

	go startSomeGoroutines(orderCh, applicationCh, reviewCh)

	for {
		select {
		case <-orderCh:
			fmt.Println("Order is arrived. Processing order...")
		case <-applicationCh:
			fmt.Println("Application is arrived. Processing application...")
		case <-reviewCh:
			fmt.Println("Review is arrived. Processing review...")
		default:
			fmt.Println("Drink tea")
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func startSomeGoroutines(first, second, third chan bool) {
	for {
		workTime := rand.Intn(5)
		someFunc := func(timeToWait int, channel chan bool) {
			time.Sleep(time.Duration(timeToWait) * time.Second)
			channel <- true
		}

		go someFunc(workTime, first)

		workTime = rand.Intn(5)

		go someFunc(workTime, second)

		workTime = rand.Intn(5)

		go someFunc(workTime, third)

		time.Sleep(time.Second)
	}
}
