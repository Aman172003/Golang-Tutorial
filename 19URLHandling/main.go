package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://bookcamp.onrender.com:5000/login?username=aman17&email=aman@gmail.com"

func main() {
	fmt.Println(myurl)

	//parsing
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("Type of query params are: %T\n", qparams)

	fmt.Println(qparams["username"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	//creating a url
	partsOfURL := &url.URL{
		Scheme:  "https",
		Host:    "bookcamp.onrender.com",
		Path:    "/login",
		RawPath: "username=nafees",
	}
	anotherURL := partsOfURL.String()
	fmt.Println(anotherURL)
}
