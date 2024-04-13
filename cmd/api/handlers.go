package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/BrentGrammer/webservice/internal/data"
)

// attach this method to the application struct type we created in main.go
func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if (r.Method != http.MethodGet) {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	    return
	}

	// NOTE: json.Marshal() will convert all data to string since we are declaring the values as string here!
	data := map[string]string{
		"status": "available",
		"environment": app.config.env,
		"version": version,
	}

	// convert the map to JSON with Marshal
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n') // just for formatting add a new line

	// set headers
	w.Header().Set("Content-Type", "application/json") // we need to do this for json since default header is set to plain text
	// now write to the response
	w.Write(js)
}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	// accepts GET and POST requests
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Display a list of books")
	}
	
	if r.Method == http.MethodPost {
		fmt.Fprintln(w, "added a new book to the reading list")
	}
}

// create wrapper multiplexor to handle different methods of request to the endpoint
func (app *application) getUpdateDeleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	// note this is how we do this using the standard net package in Go, but there are third party packages that handle allowed methods elegantly
	switch r.Method {
		case http.MethodGet:
			app.getBook(w, r)
		case http.MethodPut:
			app.updateBook(w, r)
		case http.MethodDelete:
			app.deleteBook(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	
	}
}

func (app *application) getBook(w http.ResponseWriter, r *http.Request) {
    // get the id of the book from the url string
	id := r.URL.Path[len("/v1/books/"):] // grab the part of the path after the route
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		// if we can't parse to integer we can't continue
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Display the details of book with id: %d", idInt)
}

func (app *application) updateBook(w http.ResponseWriter, r *http.Request) {
    // get the id of the book from the url string
	id := r.URL.Path[len("/v1/books/"):] // grab the part of the path after the route
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		// if we can't parse to integer we can't continue
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Update the details of book with id: %d", idInt)
}

func (app *application) deleteBook(w http.ResponseWriter, r *http.Request) {
    // get the id of the book from the url string
	id := r.URL.Path[len("/v1/books/"):] // grab the part of the path after the route
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		// if we can't parse to integer we can't continue
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Delete the details of book with id: %d", idInt)
}

