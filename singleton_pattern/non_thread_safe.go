package main

import (
	"log"
	"sync"
	"time"
)

// This is struct for mocking the connection.
type SQLConnection struct {
	connectionUrl string
	// add other details as needed
}

// This is a variable declaration in Go.
var sqlInstance *SQLConnection

func mockConnectionNonThreadSafe(threadId int) {
	if sqlInstance == nil {
		// This is a blocking call to mimic the time it takes
		//  to create a connection in real world
		time.Sleep(time.Second)

		// This is a variable assignment in Go.
		sqlInstance = &SQLConnection{
			connectionUrl: "some connection object",
		}
		log.Println("Created connection by thread id:", threadId)
	}
}

func performConcurrentAction() {
	// This is essentially needed for waiting for the program
	// to finish its concurrent tasks before exiting the program in Go.
	var wg sync.WaitGroup
	// iterate over 10 times
	// and call the mockConnectionNonThreadSafe function
	// concurrently
	for i := 0; i < 10; i++ {
		// add 1 to the wait group
		wg.Add(1)
		go func(threadId int) {
			// defer is used to ensure that the wait group is
			// decremented after the goroutine completes
			// this is done to ensure that the program doesn't
			// exit before all the goroutines complete
			defer wg.Done()
			log.Println("thread id:", threadId)
			mockConnectionNonThreadSafe(threadId)
		}(i)
	}
	// wait for all the goroutines to complete
	wg.Wait()
}

func performSequentialAction() {
	// iterate over 10 times
	// and call the mockConnectionNonThreadSafe function
	for i := 0; i < 10; i++ {
		mockConnectionNonThreadSafe(i)
	}
}

// main is the entry point for the application.
func main() {
	// below function call is for concurrent execution
	// performConcurrentAction()
	// below function call is used to execute the sequential action
	performSequentialAction()
}
