package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) GetQuarter() int {
	// need to implement
	return 0
}

type Line struct {
	Start  Point
	Finish Point
}

func (l *Line) GetLength() float64 {
	return math.Sqrt(math.Pow(l.Finish.X-l.Start.X, 2) + math.Pow(l.Finish.Y-l.Start.Y, 2))
}

type Square struct {
	Side Line
}

func (s *Square) GetPerimeter() float64 {
	return 4 * s.Side.GetLength()
}

func (s *Square) GetArea() float64 {
	return math.Pow(s.Side.GetLength(), 2)
}

func main() {
	var square Square
	var x, y float64

	fmt.Println("Enter two points that represent a side of a square in 2D plot:")
	fmt.Print("X1, Y1: ")
	fmt.Scan(&x, &y)

	square.Side.Start.X = x
	square.Side.Start.Y = y

	fmt.Print("X2, Y2: ")
	fmt.Scan(&x, &y)

	square.Side.Finish.X = x
	square.Side.Finish.Y = y

	fmt.Println(square)
	fmt.Println("The perimeter of the square, that have a side with the given line:", square.GetPerimeter())
	fmt.Println("And it's area:", square.GetArea())

	fmt.Printf("The type of square object: %T\n", square)
	fmt.Printf("The type of Side field of square object: %T\n", square.Side)
	fmt.Printf("The type of Start field of Side field of square object: %T\n", square.Side.Start)
	fmt.Printf("The type of X field of Start field of Side field of square object: %T\n", square.Side.Start.X)
}
