package main

import "fmt"

func main() {
	fmt.Print("Enter the number of tasks you've done: ")
	var numberOfTasks int
	fmt.Scan(&numberOfTasks)

	fmt.Println("Enter scores you got for each", numberOfTasks, "tasks")
	var score, sum float64
	for i := 1; i <= numberOfTasks; i++ {
		fmt.Print(i, " task: ")
		fmt.Scan(&score)
		sum += score
	}

	fmt.Println("The average score for all your tasks is", sum/float64(numberOfTasks))
}
