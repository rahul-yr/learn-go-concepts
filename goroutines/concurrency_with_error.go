package main

import (
	"log"
	"time"
)

// addTwoNumbers adds two numbers together and prints the result
func addTwoNumbers(a int, b int) {
	// sleep for a second to simulate a long running operation
	time.Sleep(2 * time.Second)
	log.Println("Result:", a+b)
}

// main is the entry point for the program
func main() {
	// call addTwoNumbers with two numbers
	go addTwoNumbers(1, 2)
	log.Println("Done")
}
