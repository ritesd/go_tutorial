package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Max Tysan")
	fmt.Println(message)
}
