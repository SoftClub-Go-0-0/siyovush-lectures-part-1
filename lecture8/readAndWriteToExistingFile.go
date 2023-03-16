package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var filename string
	fmt.Scan(&filename)
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	// open a file and truncate it
	defer file.Close()

	data := make([]byte, 999999)

	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	file.Close()

	file, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteString("new text\n")
	if err != nil {
		log.Fatal(err)
	}
}
