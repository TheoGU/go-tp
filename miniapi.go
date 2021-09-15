package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// HELPERS
func entriesList() []string {
	fmt.Fprintf("Hello world")
	raw, err := os.ReadFile("data.txt")

	if err != nil {
		return nil, err
	}
	data := strings.Split(string(raw), "\n")

	return data
}

// GET 
func helloHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		response(http.StatusMethodNotAllowed, w, req.Method+" bad request.")
	} else {
		fmt.Fprintf(w, "Hello world")
	}
}
func index(w http.ResponseWriter, req *http.Request) {
	formatTime := time.Now()
	result := fmt.Sprintf("%dh%d", formatTime.Hour(), formatTime.Minute())
	fmt.Fprintf(w, result)
}

func entries(w http.ResponseWriter, req *http.Request) {
	entries := entriesList()
	for _, rawEntry := range entries {
		entry := strings.Split(rawEntry, ":")
		fmt.Fprintf(w, entry[1]+"\n")
	}
}


// POST
func add(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	author := req.Form.Get("author")
	message := req.Form.Get("message")
	saveData, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer saveData.Close()

	if err == nil {
		fmt.Fprintf(w, "%s:%s\n", author, entry)
	}


	fmt.Fprintf(w, author+":"+message)
}

// ROUTE && SERVER
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/add", add)
	http.HandleFunc("/entries", entries)

	http.ListenAndServe(":4567", nil)
}