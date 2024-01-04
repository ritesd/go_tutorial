package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set property of the predefined Logger, including
	// the log entry to prefix and a flage to disable the printing
	// the time, source, file and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("Gladys")

	// if error was return then print the error and exit the program
	if err != nil {
		log.Fatal(err)
	}
	// if no error was return then print the greeting message
	//message, err := greetings.Hello("Max Tysan")
	fmt.Println(message)
}
