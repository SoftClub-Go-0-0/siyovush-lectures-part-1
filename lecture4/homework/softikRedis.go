package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commands := []string{
		"1. Add a value under a specified key (add)",
		"2. Get a value of a specified key (get)",
		"3. Delete a value under a specified key (delete)",
		"4. Get all the key-value pairs (get-all)",
		"5. Remove all the key-value pairs from storage (clear)",
		"6. Quit (quit)",
	}
	storage := make(map[string]string)

	fmt.Println("Salam! I'm Redis, SoftClub edition :)")
	fmt.Println("I can store any data by a specific key, like Redis. " +
		"Unfortunately for now, I can store any data in string format. " +
		"But I'm gonna be better :)")
	fmt.Println()

	fmt.Println("Here are the list of my abilities:")
	for _, command := range commands {
		fmt.Println(command)
	}

	var command, key, value string
	var exists bool

	stdin := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nType a command, you want to run:")
		fmt.Scan(&command)

		if command == "quit" {
			fmt.Println("Goodbye!")
			break
		}

		fmt.Println()

		switch command {
		case "add":
			fmt.Println("Okay, let's add something to storage.")
			fmt.Println("Enter a key:")
			stdin.Scan()
			key = stdin.Text()
			fmt.Println("Good. Now enter the value you want to save:")
			stdin.Scan()
			value = stdin.Text()
			storage[key] = value
			fmt.Println("Nice. I've added this data to storage. You can check it by getting your key.")
		case "get":
			fmt.Println("No problem, enter the key, you want to get from storage:")
			stdin.Scan()
			key = stdin.Text()
			value, exists = storage[key]
			if !exists {
				fmt.Println("Oh, seems like there is no such key in storage. Sorry.")
				break
			}
			fmt.Println("Here it is:")
			fmt.Println(value)
		case "delete":
			fmt.Println("Okay, I'll delete what you want, but be careful.")
			fmt.Println("Enter a key, you want to delete:")
			stdin.Scan()
			key = stdin.Text()
			delete(storage, key)
			fmt.Println("Done! I've deleted it.")
		case "get-all":
			fmt.Println("Here is the list of all the key-value pairs in the storage:")
			for k, v := range storage {
				fmt.Printf("%s: %s\n", k, v)
			}
		case "clear":
			fmt.Println("Removing all the data from the storage.")
			storage = make(map[string]string)
			fmt.Println("Done! Storage is cleared!")
		default:
			fmt.Println("Bro, I have no this command on the list!")
		}
	}
}
