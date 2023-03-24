package main

import (
	"fmt"
)

type Car struct {
	a string //""
	b string //""
	c int    //0
}

func main() {
	ch := make(chan Car)

	for i := 0; i < 10; i++ {
		go func(number int) {
			ch <- Car{
				a: "some color",
				b: "some car",
				c: number,
			}
		}(i)
	}

	for k := 1; k <= 10; k++ {
		fmt.Println(k, <-ch)
	}
}
