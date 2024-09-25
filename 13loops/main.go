package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Println(index, " ", day)
	}

	val := 1

	for val < 10 {
		if val == 2 {
			val++
			goto lco
		}
		if val == 6 {
			val++
			continue
		}
		if val == 8 {
			break
		}
		fmt.Println(val)
		val++
	}
lco:
	fmt.Println("Hey what's upp")
}
