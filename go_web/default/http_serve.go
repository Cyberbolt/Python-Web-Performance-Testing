package main

import (
	"encoding/json"
	"net/http"
)

type Profile struct {
	Code int `json:"code"`
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8000", nil)
}
func foo(w http.ResponseWriter, r *http.Request) {
	profile := Profile{1}
	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
