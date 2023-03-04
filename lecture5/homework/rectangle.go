package main

import "fmt"

type Rectangle struct {
	A, B float64
}

func (r *Rectangle) GetPerimeter() float64 {
	return 2 * (r.A + r.B)
}

func (r *Rectangle) GetArea() float64 {
	return r.A * r.B
}

func main() {
	var rectangle Rectangle
	fmt.Scan(&rectangle.A)
	fmt.Scan(&rectangle.B)
	fmt.Println(rectangle.GetPerimeter())
	fmt.Println(rectangle.GetArea())
}
