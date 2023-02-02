package main

import "fmt"

func main() {
	fmt.Printf("%.2f", 3.6785)
	fmt.Println()

	fmt.Printf("%T", 3.6785)
	fmt.Println()

	// \t - is tab character
	fmt.Printf("%s:\t%d", "Decimal", 11)
	fmt.Println()

	fmt.Printf("%s:\t%b", "Binary", 11)
	fmt.Println()

	// \n - is newline character
	fmt.Printf("%s:\t%o\n", "Octal", 11)

	fmt.Printf("%s:\t%x\n", "Hexadecimal", 11)

	//var s uintptr
	//s = 2453456756543236523
	//s = 18900000000000000000
	//fmt.Printf("%T:\t%v\n", s, s)
}
