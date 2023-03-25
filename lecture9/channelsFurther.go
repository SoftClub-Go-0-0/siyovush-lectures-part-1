package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// example #1
	// using buffered channel with capacity of 3 elements
	//ch := make(chan int, 1)
	//
	//ch <- 1
	//fmt.Println(<-ch)

	// example #2
	// channel blocks sending goroutine if it is full and receiving goroutine if it is empty
	//ch := make(chan int, 3)
	//
	//ch <- 1
	//ch <- 2
	//ch <- 10
	//
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)

	// example #3
	// using cap() and len() builtin functions
	//ch := make(chan int, 3)
	//
	//fmt.Println(cap(ch), len(ch))
	//ch <- 1
	//fmt.Println(cap(ch), len(ch))
	//ch <- 2
	//fmt.Println(cap(ch), len(ch))
	//ch <- 3
	//fmt.Println(cap(ch), len(ch))
	//
	//fmt.Println()
	//
	//<-ch
	//fmt.Println(cap(ch), len(ch))
	//<-ch
	//fmt.Println(cap(ch), len(ch))
	//<-ch
	//fmt.Println(cap(ch), len(ch))

	orderList := make(chan Order, 10)
	for i := 1; i <= 10; i++ {
		go AddOrder(orderList, Order{i, []string{}, rand.Intn(100) + 1})
	}

	for i := 0; i < cap(orderList); i++ {
		fmt.Println(<-orderList)
	}
}

type Order struct {
	CustomerID int
	Products   []string
	TotalCost  int
}

func AddOrder(list chan Order, order Order) {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	list <- order
}
