package main

import (
  "net/http"
  "fmt"
  "io/ioutil"
//  "os"
  "time"
)

const REQUESTS int = 1000000
const WORKERS int = 512 // blows up if this is too big. e.g.: 1024 (this is just the go channel size, not technically the number of "workers")

type HttpResponse struct {
  response *http.Response
  err error
}

type DumbResponse struct {
  value int
}

// provides a future-like http GET request function
// users can take from the returned channel to get the return value
func asyncHttpGet(url string) chan *HttpResponse {
  ch := make(chan *HttpResponse)
  go func() {
    response, err := http.Get(url)
    if err != nil {
      fmt.Print("err => ", err, "\n")
    }
    ioutil.ReadAll(response.Body)
    response.Body.Close()
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
  start_time := time.Now()
  http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 1024
  work_chan := make(chan (chan *HttpResponse), WORKERS)
  go RunRequests(work_chan, "http://localhost:6060")
  result_map := make(map[int]int)
  for http_response_chan := range work_chan {
    http_response := <-http_response_chan
    old_result, ok := result_map[http_response.response.StatusCode]
    if !ok {
      old_result = 0
    }
    result_map[http_response.response.StatusCode] = old_result + 1
  }
  duration := time.Now().Sub(start_time)

  fmt.Print("Took time: ", duration, "\n")
  fmt.Print("Req/Sec: ", float64(REQUESTS)/duration.Seconds(), "\n")
  fmt.Print("result_map => ", result_map, "\n")
}
