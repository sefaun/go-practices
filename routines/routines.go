package main

import (
	"fmt"
)

func main() {
	payload := make(chan string)

	go Foo(payload, "code:1234")
	go Bar(payload)

	var userInput string
	fmt.Scanln(&userInput)
	fmt.Println("All is well")
}

func Foo(channel chan string, content string) {
	//time.Sleep(time.Second * 3)
	fmt.Println("Foo...")
	channel <- content
}

func Bar(channel chan string) {
	fmt.Println("Bar...")
	ctx := <-channel
	fmt.Println(ctx)
}
