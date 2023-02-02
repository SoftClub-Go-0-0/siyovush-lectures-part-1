package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Signed Integers")
	fmt.Printf("int8:\tfrom %d to %d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("int16:\tfrom %d to %d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32:\tfrom %d to %d\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("int64:\tfrom %d to %d\n", math.MinInt64, math.MaxInt64)

	fmt.Println("\nUnsigned Integers")
	fmt.Printf("uint8:\tfrom %d to %d\n", 0, math.MaxUint8)
	fmt.Printf("uint16:\tfrom %d to %d\n", 0, math.MaxUint16)
	fmt.Printf("uint32:\tfrom %d to %d\n", 0, math.MaxUint32)
	fmt.Printf("uint64:\tfrom %d to %d\n", 0, math.MaxUint64-uint64(1))
}
