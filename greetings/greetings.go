package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greetin for the named person

func Hello(name string) (string, error) {
	// if no name was given, return an error with a message
	if name == "" {
		return name, errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

// at the end of return we are sending two values as return type is expected to be string, error

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {
	// in the parameter []string denot the array of string and in response map[string]string is
	// hashmap where key and value both are string
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
