package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the provided name
func Hello(name string) (string, error){
	// if no name is provided return an error
	if name == ""{
		return "", errors.New("no name provided")
	}

	// returns a greeting that embeds the provided name into the message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error){
	// A map to associate names with messages
	messages := make(map[string]string)

	// Loop through the received slice of names and calling the hello function to
	// gather a message for each value
	for _, name := range names{
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}

	return messages, nil
}

// init sets initial values
func init(){
	// This will be the function that is used to create a random number
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of the set of greeting messages the
// returned message is selected at random
func randomFormat() string{
	// A slice of message formats
	formats := []string{ // Initialize a slice of messages inline, 'slice' changes size based on remove or added content
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by
	// specifying a random index fro the slice of formats
	return formats[rand.Intn(len(formats))]
}