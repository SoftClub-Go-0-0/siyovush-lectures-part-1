package main

import "fmt"

func main() {
	var n, K, L int
	fmt.Scan(&n)
	arr := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var sum float64
	fmt.Scan(&K, &L)
	for i := K - 1; i <= L-1; i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}
