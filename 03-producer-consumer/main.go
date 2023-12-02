package main

import (
	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, totalPizzas int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch // nil if channels is closed successfully, and error is made available if not closed successfully
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making

	// run forever or until we receive a quit notification from the quit channel in producer type

	// try to make pizzas
	for {
		// try to make a pizza
		// desicion structure (selection statement)
	}
}

func main() {
	// seed the random number generator
	// As of Go 1.20, there is no reason to call Seed with a random value.
	// rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("__________________________________")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer

	// print out the ending message

}
