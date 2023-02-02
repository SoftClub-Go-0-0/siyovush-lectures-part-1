package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Floating-point numbers. Minimum value above zero to maximum value")
	fmt.Printf("float32:\tfrom %f to %f\n", math.SmallestNonzeroFloat32, math.MaxFloat32)
	fmt.Printf("float64:\tfrom %f to %f\n", math.SmallestNonzeroFloat64, math.MaxFloat64)

	fmt.Println("\nIn scientific notation:")
	fmt.Printf("float32:\tfrom %e to %e\n", math.SmallestNonzeroFloat32, math.MaxFloat32)
	fmt.Printf("float64:\tfrom %e to %e\n", math.SmallestNonzeroFloat64, math.MaxFloat64)
}
