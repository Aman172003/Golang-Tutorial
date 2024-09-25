package main

import "fmt"

func main() {
	fmt.Println("Welcome to if-else in golang")
	loginCount := 15
	var result string
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount < 20 && loginCount >= 10 {
		result = "Admin"
	} else {
		result = "Special User"
	}
	fmt.Println(result)

	// initialoxing and checking at the same time
	if num := 3; num < 10 {
		fmt.Println("Hello")
	}
}
