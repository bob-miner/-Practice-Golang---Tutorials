package main

import (
	"fmt"
	"example.com/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Gladys", "Samantha", "Darrin"}
	message, error := greetings.Hellos(names)
	
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(message)
}
