package main

import (
	"flag"
	"fmt"
)

func main() {
	//var username, password string
	//fmt.Scan(&username, &password)

	username := flag.String("username", "", "Enter your username")
	password := flag.String("password", "", "Enter your password")

	flag.Parse()

	//"developer"
	if *username != "admin" && *password != "admin" {
		fmt.Println("Access denied")
		return
	}

	fmt.Println("Access granted")

}
