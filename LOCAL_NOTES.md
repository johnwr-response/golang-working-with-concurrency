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
### 
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