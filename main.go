package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hey Welcome to my page !!")
	})
	// Hi Cosign !!
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi Cosign !!")
	})
	fmt.Println("app running on port: ", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
