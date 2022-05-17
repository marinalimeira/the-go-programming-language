package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// os.Args[0] is /var/folders[...] -- temporary place for the command execution (I think)
	firstArg := os.Args[1]
	remainingArgs := os.Args[2:]

	// Exercise 1.1
	fmt.Println("First argument: " + firstArg)

	// Exercise 1.2
	for i, arg := range remainingArgs {
		fmt.Println("Argument number " + strconv.Itoa(i+1) + ": " + arg)
	}
}
