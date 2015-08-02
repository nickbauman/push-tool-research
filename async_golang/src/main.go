package main

import (
	"bench"
	"fmt"
	"os"
	"strconv"
  "time"
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

        thenTime := time.Now()
        thenNanos := thenTime.UnixNano()
        thenMillis := thenNanos / 1000000

				bench.RunClient(secondArg, iterations)

        nowTime := time.Now()
        nowNanos := nowTime.UnixNano()
        nowMillis := nowNanos / 1000000

        duration := nowMillis - thenMillis
        req_sec := iterations / (duration / 1000.0)

        fmt.Printf("Completed: Took %d ms, %d req/sec\n", duration, req_sec)
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
