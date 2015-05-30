package main

import "fmt"

func stuff(argument int) int {
	logger.Trace("This is a trace log entry")
	fmt.Println("I received as input: ", argument, "And life is good")
	return (argument)
}
