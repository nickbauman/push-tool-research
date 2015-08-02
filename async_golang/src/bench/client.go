package bench

import (
	"fmt"
  "bytes"
	"io/ioutil"
	"net/http"
)

func RunClient(host_and_port string, num_iterations int64) {
	//results := make(map[int]int)
  var buffer bytes.Buffer
  buffer.WriteString("http://")
  buffer.WriteString(host_and_port)
  url := buffer.String()

  fmt.Printf("%s\n", url)

	resp, err := http.Get(url)
  fmt.Printf("err: %+v\n", err)
  fmt.Printf("resp: %+v\n", resp)

	if err != nil {
		// do something
    fmt.Printf("error: %+v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
  fmt.Printf("err: %+v\n", err)
  fmt.Printf("body: '%s'\n", body)
}
