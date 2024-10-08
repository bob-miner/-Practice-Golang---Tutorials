package main

import (
	"fmt"
	"example.com/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, error := greetings.Hello("")
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(message)
}
