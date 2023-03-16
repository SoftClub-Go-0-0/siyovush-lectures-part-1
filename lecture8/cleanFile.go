package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var filename string
	fmt.Scan(&filename)

	// open a file and truncate it
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}
