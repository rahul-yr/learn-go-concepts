package main

import "log"

// This is a type declaration in Go similar to enum in other languages.
// MyLocalAddress can be any string value of your choice.
// This datatype could be anything like string, int, float, etc.
type MyWebsite string

// main is a function in Go.
func main() {
	// This is a variable declaration in Go.
	var site1 MyWebsite
	// This is a variable assignment in Go.
	site1 = "rahuldev.in"
	// This is a print statement in Go.
	log.Println(site1)
}
