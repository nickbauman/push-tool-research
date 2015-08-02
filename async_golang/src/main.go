package main

import (
	"bench"
	"fmt"
	"os"
  "strconv"
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
        iterations, _ := strconv.ParseInt(thirdArg, 0, 64)
        fmt.Printf("running client against %s:%s\n", secondArg, thirdArg)
				bench.RunClient(secondArg, iterations)
			}
		} else if firstArg == "server" {
			fmt.Printf("running server on %s\n", secondArg)
      bench.RunServer(secondArg)
		} else {
			useage()
		}
	} else {
		useage()
	}
}
