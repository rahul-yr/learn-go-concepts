package main

import (
	"log"
	"sync"
	"time"
)

// wg is a wait group
// it is used to wait for all go routines to finish
var wg sync.WaitGroup

// addTwoNumbers adds two numbers together and prints the result
func addTwoNumbers(a int, b int) {
	// update the wait group counter to indicate that the go routine is done
	defer wg.Done()

	// sleep for a second to simulate a long running operation
	time.Sleep(2 * time.Second)

	log.Println("Result:", a+b)
}

// main is the entry point for the program
func main() {
	// add as many go routines as you want to run concurrently
	// and then wait for them to finish
	wg.Add(1)
	// call addTwoNumbers with two numbers
	go addTwoNumbers(1, 2)
	// wait for all go routines to finish
	wg.Wait()
	log.Println("Done")
}
