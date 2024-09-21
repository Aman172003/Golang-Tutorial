package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	// Type of fruitlist is []string
	fmt.Printf("Type of fruitlist is %T\n", fruitList)
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)
	// below makes a variablle of type slices of size 4
	// 4 will be the default size
	highScore := make([]int, 4)
	highScore[0] = 242
	highScore[1] = 345
	highScore[2] = 442
	highScore[3] = 542

	fmt.Println(highScore)
	//highscore size are redefined
	highScore = append(highScore, 555, 666, 777)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	// how to remove a value from slices based on index

	var courses = []string{"reactjs", "js", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
