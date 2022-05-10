// Echo1 prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	// Initialize variable as "zero value", in this example, as an empty string.
	var s, sep string

	// os.Args is a slice of strings
	// i++ is the same as i += 1 BUT j = i++ and --i is not valid
	// for loop is the only loop statement
	// for *initialization*; *condition*; *post* {}
	// *initialization* is optional and happens before the loop starts
	// *condition* is a boolean expression
	// for *condition* {} is also valid, like a while loop
	for i := 1; i < len(os.Args); i++ {
		// + concatenates values
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}



