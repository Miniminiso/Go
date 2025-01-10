package Assignments

import "fmt"

type Student struct {
	Name  string
	Marks []int
}

func main() {
	student := Student{
		Name:  "Sahithi",
		Marks: []int{10, 9, 8},
	}

	sum := 0
	for _, mark := range student.Marks {
		sum += mark
	}
	average := float64(sum) / float64(len(student.Marks))

	fmt.Printf("Student: %s\n", student.Name)
	fmt.Printf("Marks: %v\n", student.Marks)
	fmt.Printf("Average: %.2f\n", average)
}
