package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang!")
	greeter()

	result := adder(3, 5)

	resultPro, msg := proAdder(3, 7, 1, 9, 3, 5, 6)

	fmt.Println("The result is: ", result)
	fmt.Println("Pro result is: ", resultPro)
	fmt.Println("Pro msg is: ", msg)
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Pro adder has done it's work"
}

func greeter() {
	fmt.Println("Namastey from golang")
}
