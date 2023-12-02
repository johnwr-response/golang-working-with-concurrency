package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, totalPizzas int

// Producer is a type of struct that holds two channels:
// One for pizzas, with all information for a given pizza, including whether it was made successfully
// Another to handle end of processing (when we quit the channel)
type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

// PizzaOrder is a type of struct that describes a given pizza order. It has the order number, a message indicating
// what happened to the order and a boolean indication the order was successfully completed
type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

// Close is a method of closing the channel when we are done with it. (I.e. something is pushed to the quit channel)
func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch // nil if channels is closed successfully, and error is made available if not closed successfully
}

// makePizza attempts to make a pizza. We generate a random number from 1-12, and put in two cases where we can't make
// the pizza in time. Otherwise, we make the pizza without issue. To make things interesting, each pizza will take a
// different length of time to produce (some pizzas are harder than others).
func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d!\n", pizzaNumber)

		rnd := rand.Intn(12 + 1)
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		totalPizzas++
		fmt.Printf("Making pizza #%d. It will take %d seconds...\n", pizzaNumber, delay)

		// delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		return &p
	}
	return &PizzaOrder{pizzaNumber: pizzaNumber}
}

// pizzeria is a GoRoutine that runs in the background and calls makePizza to try to make one order each time it
// iterates through the for loop. It executes until it receives something on the quit channel. The quit channel does
// not receive anything until the consumer sends it (when the number of orders is greater than or equal to the
// constant NumberOfPizzas).
func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0

	// this loop will continue to execute, trying to make pizzas, until the quit channel receives something
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// We have tried to make a pizza (we sent something to the data channel)
			case pizzaMaker.data <- *currentPizza:

			// we want to quit, so send pizzaMaker.quit to the quitChan (a chan error)
			case quitChan := <-pizzaMaker.quit:
				// close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

// main something, something
func main() {
	// seed the random number generator
	// As of Go 1.20, there is no reason to call Seed with a random value.
	// rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!")
			}
		} else {
			color.Cyan("Done making pizzas...")
			color.Cyan("---------------------")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing channel!", err)
			}
		}
	}

	// print out the ending message
	color.Cyan("-----------------")
	color.Cyan("Done for the day.")

	color.Cyan("We made %d pizzas, but failed to make %d pizzas, with %d attempts in total", pizzasMade, pizzasFailed, totalPizzas)
	switch {
	case pizzasFailed > 9:
		color.Red("It was an awful day...")
	case pizzasFailed >= 6:
		color.Red("It was not a very good day...")
	case pizzasFailed >= 4:
		color.Yellow("It was an ok day...")
	case pizzasFailed >= 2:
		color.Yellow("It was a pretty good day!")
	default:
		color.Green("It was a great day!")

	}

}
