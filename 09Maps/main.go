package main

import "fmt"

func main() {
	// key is string and value is string
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("Ruby is deleted and now the map looks like: ", languages)

	for key, value := range languages {
		// fmt.Println(key, " ", value)
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
