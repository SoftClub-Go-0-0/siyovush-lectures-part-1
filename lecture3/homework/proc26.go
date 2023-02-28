package main

import "fmt"

func main() {
	var k, power_five_numbers_counter int
	for i := 0; i < 10; i++ {
		fmt.Scan(&k)
		if isPower5(k) {
			power_five_numbers_counter++
		}
	}
	fmt.Printf("\nPower of 5 numbers count: %d\n", power_five_numbers_counter)
}

func isPower5(k int) bool {
	power := 1
	for power < k {
		power *= 5
	}
	return power == k
}
