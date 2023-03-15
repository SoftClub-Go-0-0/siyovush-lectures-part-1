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

	_, err := fmt.Scan(&filename)
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var numbers []float64

	for scanner.Scan() {
		stringNumber := scanner.Text()
		number, err := strconv.ParseFloat(stringNumber, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	file.Close()

	var locMaxIndexes []int

	for i := 0; i < len(numbers); i++ {
		if i == 0 {
			if numbers[i] > numbers[i+1] {
				locMaxIndexes = append(locMaxIndexes, i+1)
			}
		} else if i == len(numbers)-1 {
			if numbers[i] > numbers[i-1] {
				locMaxIndexes = append(locMaxIndexes, i+1)
			}
		} else {
			if numbers[i] > numbers[i+1] && numbers[i] > numbers[i-1] {
				locMaxIndexes = append(locMaxIndexes, i+1)
			}
		}
	}

	fmt.Scan(&filename)

	file, err = os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	for _, index := range locMaxIndexes {
		indexString := strconv.Itoa(index)

		file.WriteString(indexString + " ")
	}

	file.Close()
}
