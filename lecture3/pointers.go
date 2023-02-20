package main

import "fmt"

func main() {
	var a = 10   // int
	var ptr = &a // address of a
	// ptr points to a
	// ptr contains the address of a

	fmt.Println("Initial values:")
	fmt.Printf("%v\t\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", ptr, ptr)     // *int - pointer to int
	fmt.Printf("%v\t\t%T\n", *ptr, *ptr) // dereferencing - разыменовывание
	fmt.Printf("%v\t%T\n", ptr, ptr)     // *int - pointer to int

	a += 5

	fmt.Println("\nAfter incrementing a directly:")
	fmt.Printf("%v\t\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", ptr, ptr)
	fmt.Printf("%v\t\t%T\n", *ptr, *ptr) // pointer is an alias of a variable

	//ptr += 5 // error: mismatched types *int and untyped int

	*ptr += 5 // assignment to dereferenced pointer // is the same thing as a += 5

	fmt.Println("\nAfter incrementing through the pointer to a:")
	fmt.Printf("%v\t\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", ptr, ptr)
	fmt.Printf("%v\t\t%T\n", *ptr, *ptr)

	secondPtr := &ptr
	**secondPtr += 5

	fmt.Println("\nAfter incrementing through the pointer to the pointer to a:")
	fmt.Printf("%v\t\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", ptr, ptr)
	fmt.Printf("%v\t\t%T\n", *ptr, *ptr)
	fmt.Printf("%v\t%T\n", secondPtr, secondPtr)
	fmt.Printf("%v\t%T\n", *secondPtr, *secondPtr)
	fmt.Printf("%v\t\t%T\n", **secondPtr, **secondPtr)

	thirdPtr := &secondPtr
	***thirdPtr += 5

	fmt.Println("\nAfter incrementing through the pointer to the pointer to the pointer to a:")
	fmt.Printf("%v\t\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", ptr, ptr)
	fmt.Printf("%v\t\t%T\n", *ptr, *ptr)
	fmt.Printf("%v\t%T\n", thirdPtr, thirdPtr)
	fmt.Printf("%v\t%T\n", *thirdPtr, *thirdPtr)
	fmt.Printf("%v\t%T\n", **thirdPtr, **thirdPtr)
	fmt.Printf("%v\t\t%T\n", ***thirdPtr, ***thirdPtr)
}
