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
### 
## Section: A Classic Problem: The Dining Philosophers
### What we'll cover in this section
### 
## Section: Channels, and another classic: The Sleeping Barber problem
### What we'll cover in this section
### 
## Section: Final Project - Building a Subscription Service
### What we'll cover in this section
### 
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