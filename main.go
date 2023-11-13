package main

import (
	"fmt"
	"os"
)

func main() {
	res, err := GetDOI(os.Args[1])
	if err != nil {
		fmt.Println("Error")
		return
	}

	fmt.Println(res.Message.Title[0])
	fmt.Println(res.Message.Abstract)
	fmt.Println(res.Message.Resource.Primary.URL)
}
