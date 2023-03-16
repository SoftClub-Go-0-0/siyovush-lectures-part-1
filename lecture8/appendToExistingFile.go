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
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Hello World\n")
	file.Close()
}
