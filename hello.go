package main

import (
	"fmt"
	"log"

	"github.com/goldingp/greetings"
)

// Gets a greeting message and prints it.
func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, helloErr := greetings.Hello("Patrick")
	if helloErr != nil {
		log.Fatal(helloErr)
	}

	fmt.Println(message)
}
