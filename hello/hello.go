package main

import (
	"fmt"
	"log"

	"github.com/kevinlights/greetings"
)

func main() {
	// singleGreeting()?
	multipleGreeting()
}

func multipleGreeting() {
	log.SetPrefix("greetings log: ")
	log.SetFlags(0)

	names := []string{"AAA", "BBB", "CCC"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}

func singleGreeting() {
	log.SetPrefix("greetings log: ")
	log.SetFlags(0)

	message, err := greetings.Hello("kevin")
	// message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
