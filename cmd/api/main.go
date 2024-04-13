package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/healthcheck", healthcheck)

	err := http.ListenAndServe(":4000", nil) // passing nil uses a default serve mux so you don't have to create one
	// NOTE: when using default ServeMux it is a global variable so it's possible to inject handlers by just using http.HandleFunc anywhere without specifying a server
	// To get around this, we can create our locally scoped serve mux ourselves.
	if err != nil {
		fmt.Println(err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", "dev")
	fmt.Fprintf(w, "version: %s\n", "1.0.0")
}