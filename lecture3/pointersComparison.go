package main

import "fmt"

func main() {
	var ptr *int

	fmt.Println("First pointer:")
	fmt.Printf("%v\t\t%T\n", ptr, ptr)
	fmt.Printf("%p\t\t%T\n", ptr, ptr)
	//fmt.Printf("%v\t\t%T\n", *ptr, *ptr)

	var b = 10
	var ptr2 = &b
	fmt.Println("\nSecond pointer:")
	fmt.Printf("%v\t%T\n", ptr2, ptr2)
	fmt.Printf("%p\t%T\n", ptr2, ptr2)
	fmt.Printf("%v\t\t%T\n", *ptr2, *ptr2)

	fmt.Println("\nPointers comparison:")
	fmt.Println("ptr == ptr2: ?")
	fmt.Printf("%v == %v: %t\n", ptr, ptr2, ptr == ptr2)

	// declaring another variable:
	var a = 10
	ptr = &a
	fmt.Println("\nFirst pointer after re-assignment:")
	fmt.Printf("%v\t%T\n", ptr, ptr)
	fmt.Printf("%v\t\t%T\n", *ptr, *ptr)

	fmt.Println("\nComparing pointers and their values:")
	fmt.Println("ptr == ptr2: ?")
	fmt.Printf("%v == %v: %t\n", ptr, ptr2, ptr == ptr2) // comparing the pointers

	fmt.Println("ptr != ptr2: ?")
	fmt.Printf("%v != %v: %t\n", ptr, ptr2, ptr != ptr2) // comparing the pointers

	// operators > and < are not defined on pointers, so these code lines are invalid:
	//fmt.Println("ptr > ptr2: ?")
	//fmt.Printf("%v > %v: %t\n", ptr, ptr2, ptr > ptr2) // comparing the pointers
	//
	//fmt.Println("ptr < ptr2: ?")
	//fmt.Printf("%v < %v: %t\n", ptr, ptr2, ptr < ptr2) // comparing the pointers

	fmt.Println("*ptr == *ptr2: ?")
	fmt.Printf("%v == %v: %t\n", *ptr, *ptr2, *ptr == *ptr2) // comparing dereferenced values

	ptr2 = &a
	fmt.Println("\nComparing pointers again:")
	fmt.Println("ptr == ptr2: ?")
	fmt.Printf("%v == %v: %t\n", ptr, ptr2, ptr == ptr2) // comparing the pointers

	fmt.Println("ptr != ptr2: ?")
	fmt.Printf("%v != %v: %t\n", ptr, ptr2, ptr != ptr2) // comparing the pointers

	//fmt.Println(*(ptr))      // this will work
	//fmt.Println(*(ptr + 1))  // this will not, because Go has no pointer arithmetic
	//fmt.Println(*(ptr2 - 5)) // this will not work too
}
