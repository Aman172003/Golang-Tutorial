package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var myptr *int
	// //Value of pointer is  <nil>
	// fmt.Println("Value of pointer is ", myptr)

	myNum := 23
	// & means referenceing
	var ptr = &myNum
	fmt.Println("Value of memory to which the pointer is pointing ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New number is: ", myNum)
}
