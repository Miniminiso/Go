package practiceanon

import "fmt"

func Studentstruct() {
	student := struct {
		Name  string
		Marks [3]int
	}{}

	fmt.Println("Enter the student Name:")
	fmt.Scanln(&student.Name)
	fmt.Println("Enter the 3 subject Marks:")
	for i := 0; i < len(student.Marks); i++ {
		fmt.Printf("Mark %d: ", i+1)
		fmt.Scanln(&student.Marks[i])
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
