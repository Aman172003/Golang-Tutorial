package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://bookcamp.onrender.com"

func main() {
	//make a request to a url using http req
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n", response)

	defer response.Body.Close()
	//it is necessary to close the response

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}
