package main

import "fmt"

func main() {
	var numberOfTasks int
	fmt.Println("How many tasks did you have?")
	fmt.Scan(&numberOfTasks)

	tasks := make(map[string]bool)

	for i := 0; i < numberOfTasks; i++ {
		var taskName string
		var isSolved bool
		fmt.Print("Enter task name: ")
		fmt.Scan(&taskName)
		fmt.Print("Did you solve it (true/false): ")
		fmt.Scan(&isSolved)
		tasks[taskName] = isSolved
	}

	var totalCorrectTasksNumber int
	for _, v := range tasks {
		if v {
			totalCorrectTasksNumber++
		}
	}
	fmt.Println("\nTotal correct tasks:", totalCorrectTasksNumber)

	fmt.Println("Correct tasks:")
	for taskName, isSolved := range tasks {
		if isSolved {
			fmt.Print(taskName, " ")
		}
	}
	fmt.Println()

	fmt.Println("Incorrect tasks:")
	for taskName, isSolved := range tasks {
		if !isSolved {
			fmt.Print(taskName, " ")
		}
	}
	fmt.Println()
}
