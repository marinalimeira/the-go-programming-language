package main

import (
	"fmt"
	"os"
)

func main() {
	// these are all equivalent:
	// s := ""           -> only used within a function
	// var s string
	// var s = ""        ->
	// var s string = "" -> redundant

	s, sep := "", ""

	// range produces two values: an index and the value of the element at that index
	// _ is the blank identifier, used when the syntax requires a variable but the program logic does not.
	for _, arg := range os.Args[1:] {
		s += sep + arg // the old value of *s* is garbage-collected. This is costly if the data is large
		sep = " "
	}

	fmt.Println(s)
}