package main

import "fmt"

// const type
// the first letter should be in uppercase to make it public
const LoginToken string = "vfsdfdzfdf"

func main() {
	// string datatype
	var username string = "Aman"
	fmt.Println(username)
	// %T is a placeholder
	fmt.Printf("Variable is of type: %T \n", username)

	//boolean datatype
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// int datatype
	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)
	// uint8 0-255 bound
	// uint16 0-255 bound
	// uint32 0-65535 bound
	// uint64 0-4294967295 bound

	// float data type
	var smallFloat float64 = 255.455552345435421234
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)
	//float32 has precision of 5
	//float64 has all its precision

	//default values and some aliases
	var anotherVar int
	// we don't get garbage value when variable is not initialized and used, instead we get 0
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T \n", anotherVar)

	//implicit type
	// even if you haven't mentioned that the variable is of which type then also it will assume it's type through the value initialized
	// all these things are done by lexer
	var website = "my website"
	fmt.Println(website)

	// no var style
	// it's only allowed to be declared like this in a method or function , outside of it, it will throw an error
	// := is said to be volorus operator
	numberofUser := 3000000
	fmt.Println(numberofUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
