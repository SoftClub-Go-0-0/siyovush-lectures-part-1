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
	square := Square{
		Side: Line{
			Start: Point{
				X: 4,
				Y: 4,
			},
			Finish: Point{
				X: 0,
				Y: 4,
			},
		},
	}

	fmt.Println(square)
	fmt.Println(square.GetPerimeter())
	fmt.Println(square.GetArea())

	fmt.Printf("%T\n", square)
	fmt.Printf("%T\n", square.Side)
	fmt.Printf("%T\n", square.Side.Start)
	fmt.Printf("%T\n", square.Side.Finish)
	fmt.Printf("%T\n", square.Side.Finish.X)
	fmt.Printf("%T\n", square.Side.Finish.Y)
}
