package bench

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func Run(host_and_ip string) {
	http.HandleFunc("/", hello)
	http.ListenAndServe(host_and_ip, nil)
}
