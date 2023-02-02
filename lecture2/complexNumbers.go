package main

import "fmt"

func main() {
	var x complex128 = complex(1, 2)
	fmt.Printf("x = %v\n", x)
	fmt.Printf("Real component = %v\n", real(x))
	fmt.Printf("Imaginary component = %v\n", imag(x))
}
