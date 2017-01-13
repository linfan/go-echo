package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type request struct {
	URL     string      `json:"url"`
	Method  string      `json:"method"`
	Headers http.Header `json:"headers"`
	Body    string      `json:"body"`
}

func handle(rw http.ResponseWriter, r *http.Request) {
	var err error
	var data []byte
	rr := &request{}
	rr.Method = r.Method
	rr.Headers = r.Header
	rr.URL = r.URL.String()
	data, err = ioutil.ReadAll(r.Body)
	rr.Body = string(data)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rrb, err := json.Marshal(rr)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(rrb)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(rrb)
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8000", nil)
}
