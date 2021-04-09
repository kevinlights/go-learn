package main

import (
	"fmt"
	"log"

	"github.com/kevinlights/greetings"
)

func main() {
	log.SetPrefix("greetings log: ")
	log.SetFlags(0)

	// message := greetings.Hello("kevin")
	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
