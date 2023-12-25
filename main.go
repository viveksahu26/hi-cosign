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

	log.Fatal(http.ListenAndServe(":3035", nil))
	fmt.Println("app running on port: ", 3035)
}
