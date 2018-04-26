package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	post := &Post{
		Title: "My Awesome GothamGo Talk",
		Body:  "I hope this is going well...",
	}

	body, err := json.Marshal(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
