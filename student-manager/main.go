package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Grades []int
}

var Students []Student

func AddStudent() {
	var student Student

	fmt.Print("Enter name of student: ")
	fmt.Scanln(&student.Name)

	fmt.Printf("Enter age of %s: ", student.Name)
	fmt.Scanln(&student.Age)

	fmt.Printf("Enter marks of %s in respective subjects (Math, English, Physics, and CS): ", student.Name)
	var grade int
	for i := 0; i < 4; i++ {
		fmt.Scanln(&grade)
		student.Grades = append(student.Grades, grade)
	}
	Students = append(Students, student)
}

func ShowStudentData() {
	for i, a := range Students {
		fmt.Printf("%d. Name: %s, Age: %d, Marks (Math, English, Physics, CS): %v\n", i+1, a.Name, a.Age, a.Grades)
	}
}

func main() {
	fmt.Println("Enter student data: ")
	for {
		var v string
		AddStudent()

		fmt.Println("To enter another student, type 'c'. To quit, type 'q': ")
		fmt.Scanln(&v)

		if v == "q" {
			break
		}
	}

	ShowStudentData()
}
