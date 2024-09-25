package main

import "fmt"

func main() {
	// no inheritance in golang; no super or parent
	fmt.Println("Structs in golang")

	aman := User{"Aman", "tej171995@gmail.com", false, 22}
	fmt.Println(aman)
	fmt.Printf("Aman details are %+v\n", aman)
	fmt.Println(aman.Name, " ", aman.Email, " ", aman.Status, " ", aman.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
