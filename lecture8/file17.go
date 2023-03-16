package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var filename string
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var numbers []int

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	var seriesLengths []int
	for i := 0; i < len(numbers); i++ {
		seriesLen := 1
		for i < len(numbers)-1 && numbers[i] == numbers[i+1] {
			i++
			seriesLen++
		}
		seriesLengths = append(seriesLengths, seriesLen)
	}

	fmt.Println(seriesLengths)

	// you need to write seriesLengths slice to a file
}
