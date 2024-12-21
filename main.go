package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	test2()
	test3()
	test4()
	test5()
}
func test3() {
	num := 40
	if num%2 == 0 {
		fmt.Println("Number is", num, "even")
	} else {
		fmt.Println("Number is", num, "odd")

	}
}
func test4() {

	if age := 30; age < 5 { // variable declaring inside the if function, so only if will work
		fmt.Println("Ticket is  Free")
	} else if age <= 5 || age <= 22 {
		fmt.Println("Ticket is 10")

	} else {
		fmt.Println("Ticket price is 15")
	}
	//fmt.Println(age) // it say undefined age because age value is inside the "if" statement
}
func test5() {
	marks := 50 //variable declaring outthe function
	var grade string
	if marks < 30 {
		grade = "fail"
	} else if marks >= 30 && marks <= 50 {
		grade = "pass"
	} else {
		grade = "distint"
	}

	if marks < 60 {
		fmt.Println("send message to parents")
	} else {
		fmt.Println("no message needed")
	}
	fmt.Printf("grade is %T %s\n", grade, grade)

}

func test2() {
	array1 := make([]string, 8)
	array1[0] = "sahi"
	array1[1] = "swetha"
	array1[3] = "harini"
	fmt.Println(array1)
}
