package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var firstFileName, secondFileName string
	fmt.Scan(&firstFileName, &secondFileName)

	firstFileContents, err := os.ReadFile(firstFileName) // the result is a slice of bytes
	if err != nil {
		log.Fatal(err)
	}

	secondFileContents, err := os.ReadFile(secondFileName) // the result is a slice of bytes
	if err != nil {
		log.Fatal(err)
	}

	firstFile, err := os.OpenFile(firstFileName, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	firstFile.Write(secondFileContents) // writing into the first file the content of the second file
	firstFile.Close()

	secondFile, err := os.OpenFile(secondFileName, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	secondFile.Write(firstFileContents) // writing into the second file the contents of the first file
	secondFile.Close()
}
