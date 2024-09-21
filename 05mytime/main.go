package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to timestudy of golang")
	//get the present time
	presentTime := time.Now()
	// format - 2024-09-19 14:49:26.699894 +0530 IST m=+0.000358751
	fmt.Println(presentTime)
	// 09-19-2024 14:53:39 Thursday
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// create a time
	createdDate := time.Date(2026, time.January, 17, 17, 23, 0, 0, time.UTC)
	// 2026-01-17 17:23:00 +0000 UTC
	fmt.Println(createdDate)
	// 01-17-2026 17:23:00 Saturday
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}

// GOOSE="windows" go build to build the application
