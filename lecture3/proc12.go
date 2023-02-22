package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	sortInc3(&a, &b, &c)

	fmt.Println(a, b, c)
}

func sortInc3(x, y, z *float64) {
	if *x > *y {
		*x, *y = *y, *x // (x y) min ->x max -> y
	}

	if *x > *z {
		*x, *z = *z, *x // (x z) min ->x max -> z
	}

	if *y > *z {
		*y, *z = *z, *y // (y z) min ->y max -> z
	}
}
