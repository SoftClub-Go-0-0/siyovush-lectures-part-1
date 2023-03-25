package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		fmt.Println("Hello World!")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 5)
		fmt.Println("Привет Мир!")
	}()

	wg.Wait()
	fmt.Println("Finishing program...")
}
