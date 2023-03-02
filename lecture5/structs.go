package main

import "fmt"

type Racer struct {
	Name   string
	Age    int
	Height float64
	Weight float64
	IsBald bool
}

type Student struct {
	Name        string
	Surname     string
	Hobbies     []string
	Number      string
	Performance map[string]int
}

func (s *Student) AverageScore() float64 {
	sum := 0
	for _, score := range s.Performance {
		sum += score
	}
	return float64(sum) / float64(len(s.Performance))
}

func averageScore(subjectsAndScores map[string]int) float64 {
	sum := 0
	for _, score := range subjectsAndScores {
		sum += score
	}
	return float64(sum) / float64(len(subjectsAndScores))
}

func main() {
	var object1 Student

	object1.Name = "Behruz"
	object1.Surname = "Shodiev"
	object1.Hobbies = []string{"football", "cs1.6", "sleep"}
	object1.Number = "+9929999999999"
	object1.Performance = map[string]int{
		"math":        3,
		"chem":        2,
		"fizra":       5,
		"programming": 6,
	}

	fmt.Println(object1.Name, averageScore(object1.Performance))
	fmt.Println(object1.Name, object1.AverageScore())

	var object2 Student

	object2.Name = "Muhammad"
	object2.Surname = "Hoshimov"
	object2.Hobbies = []string{"football", "cs1.6", "eating"}
	object2.Number = "+9929999999999"
	object2.Performance = map[string]int{
		"math":        3,
		"chem":        5,
		"fizra":       5,
		"programming": 6,
		"english":     4,
		"russian":     4,
	}

	fmt.Println(object2.Name, averageScore(object2.Performance))
	fmt.Println(object2.Name, object2.AverageScore())

	//fmt.Print("Name: ")
	//fmt.Scan(&student.Name)
	//fmt.Print("Surname: ")
	//fmt.Scan(&student.Surname)
	//fmt.Print("Hobbies: ")

	//var hobby string
	//for hobby != "0" {
	//	fmt.Scan(&hobby)
	//	student.Hobbies = append(student.Hobbies, hobby)
	//}

	//fmt.Print("Phone number: ")
	//fmt.Scan(&student.Number)
	//
	//var subject string
	//var score int
	//
	//student.Performance = make(map[string]int)
	//
	//for subject != "0" {
	//	fmt.Print("Subject: ")
	//	fmt.Scan(&subject)
	//	fmt.Print("Score: ")
	//	fmt.Scan(&score)
	//	student.Performance[subject] = score
	//}

	//fmt.Println(student)

	//students := make([]Student, 5)
	//for i := 0; i < 5; i++ {
	//	students[i].Name = "Student " + strconv.Itoa(i)
	//}
	//
	//for _, student := range students {
	//	fmt.Println(student)
	//}
	//
	//search := "Student 5"
	//fmt.Println()
	//fmt.Println()
	//fmt.Println("Search result")
	//
	//for _, student := range students {
	//	if student.Name == search {
	//		fmt.Println(student)
	//		break
	//	}
	//}
}
