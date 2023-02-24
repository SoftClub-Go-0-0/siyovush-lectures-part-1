package main

import (
	"fmt"
	"math"
)

func main() {
	// proc1:
	// var a, b float64
	// for i := 0; i < 5; i++ {
	// 	fmt.Scan(&a)
	// 	PowerA3(a, &b)
	// 	fmt.Println(b)
	// }

	// proc2
	// var a, a2, a3, a4 float64
	// for i := 0; i < 5; i++ {
	// 	fmt.Scan(&a)
	// 	PowerA234(a, &a2, &a3, &a4)
	// 	fmt.Println(a2, a3, a4)
	// }

	// proc3
	// var a, b, c, d float64
	// fmt.Scan(&a, &b, &c, &d)
	// var arifmetic, geometric float64
	
	// Mean(a, b, &arifmetic, &geometric)
	// fmt.Println(arifmetic, geometric)
	
	// Mean(a, c, &arifmetic, &geometric)
	// fmt.Println(arifmetic, geometric)
	
	// Mean(a, d, &arifmetic, &geometric)
	// fmt.Println(arifmetic, geometric)

	// proc4
	// var a, p, s float64
	// for i := 0; i < 3; i++ {
	// 	fmt.Scan(&a)
	// 	TrianglePS(a, &p, &s)
	// 	fmt.Printf("%.2f\t%.2f\n", p, s)
	// }

	// proc5
	var x1, x2, y1, y2, p, s float64
	for i := 0; i < 3; i++ {
		fmt.Scan(&x1, &y1, &x2, &y2)
		RectPS(x1, y1, x2, y2, &p, &s)
		fmt.Printf("%.2f\t%.2f\n", p, s)
	}
}

func PowerA3(a float64, b *float64) {
	*b = a * a * a
}

func PowerA234(a float64, b, c, d *float64) {
	*b = a * a
	*c = *b * a
	*d = *c * a
}

func Mean(x, y float64, a, g *float64) {
	*a = (x + y) / 2
	*g = math.Sqrt(x * y)
}

func TrianglePS(a float64, P, S *float64) {
	*P = 3 * a
	*S = a * a * math.Sqrt(3) / 4
}

func RectPS(x1, y1, x2, y2 float64, P, S *float64) {
	a := math.Abs(x1 - x2)
	b := math.Abs(y1 - y2)

	*P = 2 * (a + b)
	*S = a * b
}