package bench

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func makeRequests(iterations int64, url string) <-chan *http.Response {
	out := make(chan *http.Response)
	go func() {
		for i := int64(0); i < iterations; i++ {
			resp, _ := http.Get(url)
			out <- resp
		}
		close(out)
	}()
	return out
}

func processResponses(responsesChannel <-chan *http.Response, results map[int]int) map[int]int {
	for resp := range responsesChannel {
		val, ok := results[resp.StatusCode]
		if ok {
			results[resp.StatusCode] = val + 1
		} else {
			results[resp.StatusCode] = 1
		}
		defer resp.Body.Close()
		ioutil.ReadAll(resp.Body)
	}
	return results
}

func RunClient(host_and_port string, num_iterations int64) {
	results := make(map[int]int)
	var buffer bytes.Buffer

	buffer.WriteString("http://")
	buffer.WriteString(host_and_port)
	url := buffer.String()

	fmt.Printf("running %d against %s\n", num_iterations, url)

	res := processResponses(makeRequests(num_iterations, url), results)

	fmt.Printf("results: %+v\n", res)
}
