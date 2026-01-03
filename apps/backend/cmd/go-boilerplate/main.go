package main

import (
	"fmt"
	"log"
	"net/http"
)

func welcome(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello to this welcome endpoint!"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/welcome", welcome)
	fmt.Println("Server started on Port 8080 :)")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server")
	}
}
