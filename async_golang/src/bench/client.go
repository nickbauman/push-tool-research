package bench

import (
	"fmt"
  "bytes"
	"io/ioutil"
	"net/http"
)

func RunClient(host_and_port string, num_iterations int64) {
	results := make(map[int]int)
  var buffer bytes.Buffer

  buffer.WriteString("http://")
  buffer.WriteString(host_and_port)
  url := buffer.String()

  fmt.Printf("%s\n", url)

  for i := int64(0); i<num_iterations; i++ {
		resp, _ := http.Get(url)
    //fmt.Printf("err: %+v\n", err)
    //fmt.Printf("resp: %+v\n", resp)
    val, ok := results[resp.StatusCode]
    if(ok) {
      results[resp.StatusCode] = val + 1
    } else {
      results[resp.StatusCode] = 1
    }
	  defer resp.Body.Close()
	  ioutil.ReadAll(resp.Body)
	}

  fmt.Printf("results: %+v\n", results)
}
