package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to pizza app")
	fmt.Println("Please rate btwn 1 to 10")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	// suppose 4 is take as a input to string me as a 4\n save hoga, to jab int me convert ya float me convert kr rhe ho to \n hatao
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error is ", err)

	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
