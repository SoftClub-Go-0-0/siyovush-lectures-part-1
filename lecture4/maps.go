package main

import "fmt"

func main() {
	// map[key type]value type

	// map declaration
	students := make(map[string]string)
	fmt.Printf("%v\t%T\n", students, students)

	// adding elements to map
	students["Khujaev"] = "Muhammad"
	fmt.Printf("%v\t%T\n", students, students)

	fmt.Println()

	// map declaration
	studentsScores := make(map[string]int)
	fmt.Printf("%v\t%T\n", studentsScores, studentsScores)

	studentsScores["Mehrdod"] = 944
	fmt.Printf("%v\t%T\n", studentsScores, studentsScores)

	// getting a value under specific key
	fmt.Println(students["Khujaev"])
	fmt.Println(students["Shodiev"]) // key doesn't exist
	fmt.Println(studentsScores["Mehrdod"])

	// updating a value
	students["Khujaev"] = "Muhammadjon"
	fmt.Println(students)

	// checking the existence of a key
	mehrdodsScore := studentsScores["Mehrdod"]
	fmt.Println("Mehrdod's score:", mehrdodsScore)

	var exists bool
	mehrdodsScore, exists = studentsScores["Mehrdod"]
	fmt.Println("Do we have Mehrdod's score? ", exists)
	fmt.Println("Mehrdod's score:", mehrdodsScore)

	fmt.Println()

	siyovushsScore, exists := studentsScores["Siyovush"]
	fmt.Println("Do we have Siyovush's score? ", exists)
	fmt.Println("Siyovush's score:", siyovushsScore)

	usersScores := make(map[string]int)
	_, exists = usersScores["paraparadox"]
	if exists {
		usersScores["paraparadox"] += 10
	} else {
		usersScores["paraparadox"] = 15
	}

	// deleting a value under a key
	delete(students, "Khujaev")
	delete(studentsScores, "Mehrdodjon")
	fmt.Println(students, studentsScores)
	delete(studentsScores, "Mehrdod")
	fmt.Println(studentsScores)

	// iteration through all the elements of a map
	for k, v := range usersScores {
		fmt.Printf("User %s has %d scores in total\n", k, v)
	}

	fmt.Println()

	usersScores["dubai"] = 1250
	for k, v := range usersScores {
		fmt.Printf("User %s has %d scores in total\n", k, v)
	}

	// practice with map

	products := make(map[string]float64)

	var productName string
	var productPrice float64

	for i := 0; i < 5; i++ {
		fmt.Scan(&productName, &productPrice)
		products[productName] = productPrice
	}

	for k, v := range products {
		fmt.Printf("Product %s is %.2f somoni\n", k, v)
	}
}
