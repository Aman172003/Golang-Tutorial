package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	// when is called as json it will be shown as coursename
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string
	//show nothing when tag is empty
	Tags []string `json:"tags,omitempty"`
}

func main() {
	// EncodedJson()
	DecodeJson()
}

func EncodedJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "Youtube", "abc123", []string{"web-dev", "JS"}},
		{"MERN Bootcamp", 199, "Youtube", "bcd123", []string{"fullstack", "JS"}},
		{"Angular Bootcamp", 299, "Youtube", "xyz123", nil},
	}
	// package above data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
            "coursename": "MERN Bootcamp",
            "Price": 199,
            "Platform": "Youtube",
            "Password": "bcd123",
            "Tags": ["fullstack", "JS"]
        }
	`)
	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	// some cases where you want to create key-value pair
	// value is not guaranteed of any specific type so use interface for that
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and Value is %v\n", k, v)
	}
}
