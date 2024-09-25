package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// create a file
	content := "This needs to go in a file - aman.gmail.com"
	file, err := os.Create("./mylcogofile.txt")

	checkNilError(err)

	// to  write we use io
	// below will retutn the length of the new file created or error
	length, err := io.WriteString(file, content)

	checkNilError(err)
	fmt.Println(length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

// function to read the file
func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilError(err)
	fmt.Println(databyte)
	fmt.Println(string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		// panic will shut down the execution of the program and show what the error is
		panic(err)
	}
}
