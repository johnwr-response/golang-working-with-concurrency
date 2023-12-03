package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}
func main() {
	// create two channels
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (q to quit)")
	for {
		// print a prompt
		fmt.Print("-> ")

		// get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if "q" == strings.ToLower(userInput) {
			break
		}

		ping <- userInput
		// wait for a response
		response := <-pong
		fmt.Printf("Response : %s\n", response)
	}

	fmt.Println("All done. Closing channels.")
	close(ping)
	close(pong)
}
