// we keep the handlers in the main package still
package main

import (
	"fmt"
	"net/http"
)

// attach this method to the application struct type we created in main.go
func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if (r.Method != http.MethodGet) {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	    return
	}
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", cfg.)
	fmt.Fprintf(w, "version: %s\n", version)
}

