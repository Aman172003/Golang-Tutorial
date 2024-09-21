package main

import "fmt"

func main() {

	var arr [4]string

	arr[0] = "Aman"
	arr[1] = "Shubham"
	arr[3] = "Nafees"
	//there will be some extra space after shubham because there is nothong placed at 2nd index
	// [Aman Shubham  Nafees]
	fmt.Println("All names: ", arr)
	// 4
	fmt.Println(len(arr))

	var marks = [3]int{12, 15, 11}
	fmt.Println(marks)
}
