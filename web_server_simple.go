package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Println("received a request on root path / ")
	fmt.Fprintf(w, "Hello, "+req.URL.Path[1:])
}

func handleSubmitIdea(w http.ResponseWriter, req *http.Request) {
	fmt.Println("received a request on root path /submitIdea ")
	fmt.Fprintf(w, "Idea submitted, "+req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/submitIdea", handleSubmitIdea)

	err := http.ListenAndServe("localhost:3002", nil)

	if err != nil {
		log.Fatal("ListenAndServer failed: ", err.Error())
	}
}
