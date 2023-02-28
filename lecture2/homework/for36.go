package main

import "fmt"

func main() {
	var n, k int
	fmt.Print("N = ")
	fmt.Scan(&n)
	fmt.Print("K = ")
	fmt.Scan(&k)
	sum := 0.0
	power_k := 1.0
	for i := 1; i <= n; i++ {
		power_k = float64(i)
		for j := 1; j < k; j++ {
			power_k *= float64(i)
		}
		sum += power_k
	}
	fmt.Printf("\n%.2f\n", sum)
}