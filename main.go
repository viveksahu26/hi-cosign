package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hey Welcome to my page !!, %q", html.EscapeString(r.URL.Path))
	})
	// Hello!
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi Cosign !!")
	})

	log.Fatal(http.ListenAndServe(":3035", nil))
}
