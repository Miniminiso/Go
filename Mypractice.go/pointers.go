package main

import "fmt"

func arr() *int {
	arr := [3]int{1, 2, 3}
	ptr := &arr[1]
	fmt.Println("arr original", arr)
	return ptr
}

func printarray(ptr *int) {
	fmt.Println("value of ptr", *ptr)

}
func main() {
	ptr := arr()
	printarray(ptr)

}
