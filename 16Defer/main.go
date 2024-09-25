package main

import "fmt"

func main() {
	// defer statements last me run hote hai reverse order me
	defer fmt.Println("Hello Duniya")
	defer fmt.Println("bhlll")
	fmt.Println("Hello Bhai")
	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
