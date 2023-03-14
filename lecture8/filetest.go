package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var filename string

	_, err := fmt.Scan(&filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var numbers []int
	for scanner.Scan() {
		stringNumber := scanner.Text()
		number, err := strconv.Atoi(stringNumber)
		if err != nil {
			fmt.Println(err)
			return
		}
		numbers = append(numbers, number)
	}

	file.Close()

	fmt.Printf("%v\t%T\n", file, file)
	fmt.Println(numbers)

	file, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
	for _, number := range numbers {
		_, err := file.WriteString(strconv.Itoa(number*2) + " ")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	file.Close()

	fmt.Println("All done!")

}
