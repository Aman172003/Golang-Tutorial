package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza")

	//comma ok \\ comma error syntax
	// read until /n is there
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanx for rating ", input)
	fmt.Printf("Type of this rating is %T", input)
}
