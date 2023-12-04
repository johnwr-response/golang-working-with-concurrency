# Working with Concurrency in Go (Golang)

## Section: Introduction
### Introduction
- Why and when you should and should not use concurrency
- Go's Philosophy
- Don't communicate by sharing memory, share memory by communicating
- Type the word Go before call to functions and suddenly they are running concurrently in the background
- WaitGroups
- Channels
#### Share Memory by Communicating
- Or, in more detail:  
`Don't over-engineer things by using shared memory and complicated, error-prone synchronization primitives; instead, use message-passing between GoRoutines so variables and data can be used in the appropriate sequence.`
- When to use
  - A golden rule for concurrency: if you don't need it, don't use it
  - Keep your application's complexity to an absolute minimum;it's easier to write, easier to understand, and easier to maintain.
- What We'll cover
  - We'll start by looking at the basic types in the sync package: mutexes (semaphores) and wait groups
  - We'll go through three of the classic computer science problems:
    - The Producer/Consumer problem
    - The Dining Philosopher problem
    - The Sleeping Barber problem
  - These are classic problems for a reason: they give students excellent exposure to the problems inherent in concurrent programming, and they force them to figure out an efficient method of solving such problems.
  - We'll also cover a more real-world scenario, where we build a subset of a larger (imaginary) problem where a user of a service wants to subscribe and buy one of a series of available subscriptions.
  - When they purchase a subscription, we will generate an invoice, send an email, generate a PDF manual and send that to the user as well.
  - We'll do all those things concurrently.

### A bit about me
### Installing Go
- Setup for windows:
  ```powershell
  winget install -e --id GoLang.Go
  ```

### Installing Visual Studio Code
- Setup for windows:
  ```powershell
  winget install -e --id Microsoft.VisualStudioCode
  ```
- You want to make sure all the `Go: Install/Update Tools` are enabled
- Suggestion: Also install extension: `gotemplate-syntax`

### JetBrains Toolbox (GoLand)
- Setup for windows:
  ```powershell
  winget install -e --id JetBrains.Toolbox
  ```

### Installing Make
- Setup for windows:
  ```powershell
  winget install -e --id GnuWin32.Make
  ```

### Asking for help
### Mistakes: we all make them

## Section: Goroutines, the go keyword, and WaitGroups
### What we'll cover in this section
- GoRoutines
  - Running things in the background, or concurrently
  - Incredibly simple to use
  - Create problems
  - Several ways to solve these problems

### Creating Goroutines and identifying a problem
- First create folder and init a main module
  ```powershell
  $incNum = "01"
  $folderName = "first-example"

  md $incNum-$folderName
  cd $incNum-$folderName
  go mod init $folderName
  ni main.go -type file -Value "package main`n`nfunc main() {`n`n}"
  ```
- NOTE! Every Go program has a Goroutine, the main() function itself is actually a Goroutine 
- Goroutines are in effect very light threads that are run by the Go scheduler behind the scenes
- The keyword `Go` tells the compiler to execute the code in its own Goroutine
  - In this example, the program executes and runs through so quickly that the new thread does not start in time.
  - Using sleeps to fix that should result in you looking for a new job!

### WaitGroups to the rescue
- Note! It doesn't matter which order you spun the go routines, the Go scheduler will execute them in its own order
- Wait Groups are the first and probably the easiest way of dealing with concurrency 
- Note: Defer means that whatever comes after the `defer` keyword, do not execute it until the current function exits

### Writing tests with WaitGroups
- Note! If mismatch between number added to Wait Group and actual Goroutines spawned, an error will be shown:
  - fatal error: all goroutines are asleep - deadlock!
- An important point when creating tests for functions that are expected to run in the background is to rin it in the background also in the test 

### Challenge: working with WaitGroup
- Modify the code in [challenge-1](challenge-1/main.go) to
  - call the function `updateMessage()` as goroutines
  - implement WaitGroups so that the program runs properly and prints out the different messages consistently in order
  - include tests for all three functions: `updateMessage()`, `printMessage()`, and `main()`

### Solution to Challenge

## Section: Race conditions, Mutexes, and an Introduction to Channels
### What we'll cover in this section
#### sync.Mutex
- Mutex (mutual exclusion) - allows us to deal with race conditions
- Relatively simple to use
- Dealing with shared resources and concurrent/parallel goroutines
- Lock/Unlock
- We can test for race conditions when running code, or when testing it

#### Race Conditions
- Race conditions occur when multiple GoRoutines try to access the same data
- Can be difficult to spot when reading code
- Go allows us to check for them when running a program, or when testing our code with go test 

#### Channels
- Channels are a means of having GoRoutines share data
- They can talk to each other both ways, you can have Unidirectional and Bidirectional channels
- You can have as many channels as you want
- This is Go's philosophy of having things share memory by communicating, rather than communicating by sharing memory
- Introducing a classic computer science problem: The Producer/Consumer problem

### Race Conditions: an example
- To run the race condition check mentioned, you have to install a C compiler first
  - Download [TDM-GCC](https://jmeubank.github.io/tdm-gcc/download/)
  - Restart all Powershell windows, including in IDE
- Create folder and init a main module
  ```powershell
  $incNum = "02"
  $folderName = "mutex-example"

  md $incNum-$folderName
  cd $incNum-$folderName
  go mod init $folderName
  ni main.go -type file -Value "package main`n`nfunc main() {`n`n}"
  ```
- To test for race conditions, use option `-race` when running your program 
  ```powershell
  go run -race .
  ```

### Adding sync.Mutex to our code
- DO NOT COPY A MUTEX ONCE IT HAS BEEN CREATED!

### Testing for race conditions
### A more complex example
### Writing a test for our weekly income project
- NOTE! This test sometimes just hangs on Windows computers. Just move on, this is not important!

### Producer/Consumer - Using Channels for the first time
- Introducing the [Producer-consumer](https://en.wikipedia.org/wiki/Producer%E2%80%93consumer_problem) problem (bounded-buffer problem)
- Create folder and init a main module
  ```powershell
  $incNum = "03"
  $folderName = "producer-consumer"

  md $incNum-$folderName
  cd $incNum-$folderName
  go mod init $folderName
  ni main.go -type file -Value "package main`n`nfunc main() {`n`n}"
  git add .
  ```
- This version of the Producer-consumer problem is based on a pizzeria
  - We have a pizzeria which have som people who make pizzas `The Producer`
  - We have one or more customers who place orders `The Consumer` 
  - A number of things can go wrong, and also the customers should get their Pizzas
  - The goal is to write a solution to the Producer-consumer problem
  - This problem can be quite complex as we will see here
- Introducing Channels that allows GoRoutines to exchange data to and from other GoRoutines

### Getting started with the Producer - the pizzeria function
- Introducing a simple package to use colorized outputs: [color](https://github.com/fatih/color)
  ```powershell
  go get github.com/fatih/color
  ```
- NOTE: As of Go 1.20, seed is deprecated. There is no reason to call Seed with a random value.
- THE GOLDEN RULE OF CHANNELS: When finished with it; close it!

### Making a pizza: the makePizza function
### Finishing up the Producer code
- Introducing the `select` statement:
  - A `select` statement is ***only*** useful for `channels`
  - It is very similar to a `switch` statement

### Creating and running the consumer: ordering a pizza
### Finishing up our Producer/Consumer project

## Section: A Classic Problem: The Dining Philosophers
### What we'll cover in this section
- Introducing the [The Dining Philosophers](https://en.wikipedia.org/wiki/Dining_philosophers_problem) problem

#### The Dining Philosophers
- A classic computer science problem introduced by Dijkstra in 1965
- Five philosophers live in a house together, and they always dine together at the same table, sitting in the same place
- They always eat a special kind of spaghetti which requires two forks
- There are two and only two forks next to each plate, which means that no two neighbours can be eating at the same time
- How do you write a program that ensures that no philosopher will starve

![Illustration](https://upload.wikimedia.org/wikipedia/commons/7/7b/An_illustration_of_the_dining_philosophers_problem.png)

### Getting started with the problem
- Create folder and init a main module
  ```powershell
  $incNum = "04"
  $folderName = "dining-philosophers"

  md $incNum-$folderName
  cd $incNum-$folderName
  go mod init $folderName
  ni main.go -type file -Value "package main`n`nfunc main() {`n`n}"
  git add .
  ```

### Implementing the diningProblem logic
- What makes this a fun problem to solve is the possible Logical Race Condition here. (Like is everyone takes their left fork first)
- We need to make sure that two people does not take the wrong fork.
- Logical Race Conditions are not detected by the `-race` flag 
- If we make sure that we always get a lock on the lower numbered fork first, we will never run into that problem

### Challenge: Printing out the order in which the meal is finished
- At the end, print out the order in which the diners finished the meal.

### Solution to challenge

### Writing a test for our program

## Section: Channels, and another classic: The Sleeping Barber problem
### What we'll cover in this section
- Introducing the [Sleeping Barber](https://en.wikipedia.org/wiki/Sleeping_barber_problem) problem

#### Channels
- A means of allowing communication to and from a GoRoutine
- Channels can be buffered, or unbuffered
- Once you're done with a channel, you must close it
- Channels typically only accept a given type or interface

#### The Sleeping Barber
- A classic computer science problem introduced by Dijkstra in 1965
- A barber goes to work in a barbershop with a waiting room with a fixed number of seats
- If no one is in the waiting room, the barber goes to sleep
- When a client shows up, if there are noe seats available, he or she leaves
- If there is a seat available, and the barber is sleeping, the client wakes the barber up and gets a haircut
- If the barber is busy, the client takes a seat and waits for his or her turn
- Once the shop closes, no more clients are allowed in, but the barber has to stay until everyone who is waiting gets their haircut

### Introduction to channels
- Create folder and init a main module
  ```powershell
  $incNum = "05"
  $folderName = "simple-channels"

  md $incNum-$folderName
  cd $incNum-$folderName
  go mod init $folderName
  ni main.go -type file -Value "package main`n`nfunc main() {`n`n}"
  git add .
  ```
- To mark a channel as receive and send, define the channel as `chan`
- To mark a channel as receive only, define the channel as `<-chan`
- To mark a channel as send only, define the channel as `chan<-`

### The select statement
- Create folder and init a main module
  ```powershell
  $incNum = "06"
  $folderName = "channel-select"

  md $incNum-$folderName
  cd $incNum-$folderName
  go mod init $folderName
  ni main.go -type file -Value "package main`n`nfunc main() {`n`n}"
  git add .
  ```
- If there are more than one matches to a select statement, it chooses one of them randomly

### Buffered Channels
- Create folder and init a main module
  ```powershell
  $incNum = "07"
  $folderName = "buffered-channels"

  md $incNum-$folderName
  cd $incNum-$folderName
  go mod init $folderName
  ni main.go -type file -Value "package main`n`nfunc main() {`n`n}"
  git add .
  ```
- Unbuffered channels take one message at a time
- Buffered channels take x number of messages at a time

### Getting started with the Sleeping Barber project
- Create folder and init a main module
  ```powershell
  $incNum = "08"
  $folderName = "sleeping-barber"

  md $incNum-$folderName
  cd $incNum-$folderName
  go mod init $folderName
  ni main.go -type file -Value "package main`n`nfunc main() {`n`n}"
  git add .
  ```

### Defining some variables, the barber shop, and getting started with the code
- Adding simple package to use colorized outputs: [color](https://github.com/fatih/color)
  ```powershell
  go get github.com/fatih/color
  ```

### Adding a Barber
### Starting the barbershop as a GoRoutine
### Sending clients to the shop
### Trying things out
  ```powershell
  go run .
  go run -race .
  ```
- Problem solved with no mutexes, no WaitGroups, only channels. Share memory by communicating

## Section: Final Project - Building a Subscription Service

### What we'll cover in this section

#### The final project
- A more `real-world` use of concurrency
- A fictitious service that allows people to buy a subscription
- This section sets up the web application

### Setting up a simple web application
- Create folder and init a main module
  ```powershell
  $incNum = "09"
  $folderName = "final-project"

  md $incNum-$folderName\cmd\web
  cd $incNum-$folderName
  go mod init $folderName
  cd cmd\web
  ni main.go -type file -Value "package main`n`nfunc main() {`n`n}"
  cd ..\..
  git add .
  md db-data\postgres
  md db-data\redis
  ```
- Adding some packages to use
  ```powershell
  go get github.com/jackc/pgconn
  go get github.com/jackc/pgx/v4
  go get github.com/jackc/pgx/v4/stdlib
  go get github.com/alexedwards/scs/v2
  go get github.com/alexedwards/scs/redisstore
  go get github.com/go-chi/chi/v5
  go get github.com/go-chi/chi/v5/middleware
  ```

### Setting up our Docker development environment
- Running docker-compose
  ```powershell
  docker-compose up -d 
  ```
- Also setup a database client (if needed)
  - [Beekeeper Studio (Community Edition)](https://github.com/beekeeper-studio/beekeeper-studio)

### Adding postgres
### Setting up a Makefile
- Running application with make
  ```powershell
  make start
  ```
### Adding sessions & Redis
### Setting up the application config





## Section: Sending Email Concurrently
### What we'll cover in this section
### 
## Section: Registering a User and Displaying Plans
### What we'll cover in this section
### 
## Section: Adding Concurrency to Choosing a Plan
### What we'll cover in this section
### 
## Section: Testing
### What we'll cover in this section
### 