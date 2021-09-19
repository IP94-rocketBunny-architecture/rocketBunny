package main

import (
	"encoding/json"
	// "log"
	"net/http"
	"time"
)

func main() {
	type Time struct {
		Time string `json:"time"`
	}

	http.HandleFunc("/time", func(w http.ResponseWriter, req *http.Request) {
		test := &Time{Time: time.Now().Format(time.RFC3339)} //current time
		res, _ := json.Marshal(test)
		w.Header().Set("content-type", "aplication/json")
		w.WriteHeader(200)
		w.Write(res)
	})
}
