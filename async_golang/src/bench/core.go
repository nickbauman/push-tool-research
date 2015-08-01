package main

import (
	"bench/client"
	"bench/server"
	"fmt"
	"os"
)

func useage() {
	fmt.Printf("choose either 'server' or 'client' followed by a host/ip value respectively\n")
	fmt.Printf("example: $ <binary> client localhost:8080 100000\n")
	fmt.Printf("or:      $ <binary> server localhost:8080\n")
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) >= 2 {
		firstArg := argsWithoutProg[0]
		secondArg := argsWithoutProg[1]

		if firstArg == "client" {
			if len(secondArg) > 1 {
				thirdArg := argsWithoutProg[2]
				fmt.Printf("running client against %s\n", secondArg, thirdArg)
				client.Run(secondArg, thirdArg)
			}
		} else if firstArg == "server" {
			fmt.Printf("running server on %s\n", secondArg)

		} else {
			useage()
		}
	} else {
		useage()
	}
}
