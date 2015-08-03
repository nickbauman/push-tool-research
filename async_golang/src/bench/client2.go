package main

import (
  "net/http"
  "fmt"
  "io/ioutil"
//  "os"
//  "time"
)

const REQUESTS int = 100000 // blows up right now if too big...
const WORKERS int = 1024

type HttpResponse struct {
  response *http.Response
  err error
}

// provides a future-like http GET request function
// users can take from the returned channel to get the return value
func asyncHttpGet(url string) chan *HttpResponse {
  ch := make(chan *HttpResponse)
  go func() {
    response, err := http.Get(url)
    defer response.Body.Close()
    ioutil.ReadAll(response.Body)
    ch <- &HttpResponse{response, err}
    close(ch)
  }()
  return ch
}

func RunRequests(work_chan chan (chan *HttpResponse), url string) {
  for i := 0; i < REQUESTS; i++ {
    work_chan <- asyncHttpGet(url)
  }
  close(work_chan)
}

func main() {
  work_chan := make(chan (chan *HttpResponse), WORKERS)
  go RunRequests(work_chan, "http://localhost:6060")
  result_map := make(map[int]int)
  for http_response_chan := range work_chan {
    fmt.Print("Hello!\n")
    http_response := <-http_response_chan
    old_result, ok := result_map[http_response.response.StatusCode]
    if !ok {
      old_result = 0
    }
    result_map[http_response.response.StatusCode] = old_result + 1
  }
  fmt.Print("result_map => ", result_map)
  /*for i := 0; i < 100000; i++ {
    result := <- asyncHttpGet("http://localhost:6060")
    fmt.Print("Status => ", result.response.StatusCode, "\n")
  }*/
}
