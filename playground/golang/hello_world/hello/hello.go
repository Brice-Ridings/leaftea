package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Create a slice of names
	names := []string {"Brice", "Sierra", "Lincoln", "Emilia"}

	// Get a greeting message and print it.
	message, err := greetings.Hellos(names)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(message)
}