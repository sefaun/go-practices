package main

import (
	"fmt"
	httprequest "practice/http-request"
	"practice/types"
)

func main() {

	result := make(chan types.LoginResponse)

	go httprequest.Login(result)

	for n := range result {
		fmt.Println(n.Success)
		fmt.Println(n.Data)
		fmt.Println(n.Message)
	}

}
