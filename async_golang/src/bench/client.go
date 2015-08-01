package bench

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Run(url string, num_iterations int) {
	//results := make(map[int]int)

	resp, err := http.Get(url)
	fmt.Printf("err %+v", err)
	fmt.Printf("resp %+v", resp)

	if err != nil {
		// do something
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("err %+v", err)
	fmt.Printf("body %+v", body)
}
