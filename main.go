package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Config struct {
	Version string `json:"version"`
}

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

func version(w http.ResponseWriter, req *http.Request) {
	var config Config
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Fprint(w, "Unable to retrieve version information.")
	}
	configBytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(configBytes, &config)
	fmt.Fprint(w, config.Version)

}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/version", version)

	http.ListenAndServe(":80", nil)
}
