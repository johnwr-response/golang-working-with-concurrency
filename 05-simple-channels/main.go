package main

import (
	"fmt"
	"strings"
)

// shout has twp parameters: a receive-only chan ping, and a send-only chan pong. Note the use of `<-` in function
// signature. It simply takes whatever string it gets from the ping channel, converts it to uppercase and appends a few
// exclamation marks, and then sends the transformed text to the pong channel.
func shout(ping <-chan string, pong chan<- string) {
	for {
		//read from the ping channel. Note that the GoRoutine waits here, it blocks until something is received on ping
		s, ok := <-ping
		if !ok {
			// do something
		}

		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	// create two channels. Ping is what we send to, pong is what comes back
	ping := make(chan string)
	pong := make(chan string)

	// start a GoRoutine
	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (q to quit)")
	for {
		// print a prompt
		fmt.Print("-> ")

		// get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if "q" == strings.ToLower(userInput) {
			// jump out of for loop
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
