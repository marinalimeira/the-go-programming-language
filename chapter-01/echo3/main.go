package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//fmt.Println(strings.Join(os.Args[1:], " "))

	allArgs := os.Args[1:]
	allArgsStr := strings.Join(allArgs, " ")

	fmt.Println(allArgsStr)

	fmt.Println(allArgs) // prints with []
}