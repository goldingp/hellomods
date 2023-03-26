package main

import (
	"fmt"

	"github.com/goldingp/greetings"
)

// Gets a greeting message and prints it.
func main() {
	message := greetings.Hello("Patrick")

	fmt.Println(message)
}
