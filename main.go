package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type requestBody struct {
	Task string `json:"task"`
}

var task string = "default task"

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Post Only", http.StatusMethodNotAllowed)
		return
	}

	var request requestBody

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task = request.Task

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Task updated successfully")

}

func handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Get Only", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "hello, %s", task)
}

func main() {
	http.HandleFunc("/", handleGet)
	http.HandleFunc("/task", handlePost)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
