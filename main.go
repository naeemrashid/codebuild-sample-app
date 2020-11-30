package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "<h1>Hello World!</h1><div>Serving on %s</div>", hostname)
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":80", nil)
}
