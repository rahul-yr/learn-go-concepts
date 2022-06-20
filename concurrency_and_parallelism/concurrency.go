package main

import (
	"log"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

// threadProfile is a profile of threads created by the program.
var threadProfile = pprof.Lookup("threadcreate")

func concurrencyInAction() {
	// defer the call to wg.Done().
	defer wg.Done()

	// sleep for a second.
	time.Sleep(time.Second * 2)

}

// init is called before main.
// init is predefined function in go like main(), but it is not mandatory to use it.
func init() {
	// set the number of CPUs to use.
	// By default, the number of CPUs is the number of CPUs on the machine.
	// Just to show we can change the number of CPUs
	// I have added this below command.
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// define the number of goroutines to use.
	count := 10
	// add the number of goroutines to wait for.
	wg.Add(count)
	// log the number of threads created.
	log.Println("Before thread count : ", threadProfile.Count())
	// loop through the number of goroutines.
	for i := 0; i < count; i++ {
		// call the concurrencyInAction function.
		go concurrencyInAction()
	}

	// wait for the goroutines to finish.
	wg.Wait()
	// log the number of threads created.
	log.Println("After thread count : ", threadProfile.Count())
}
