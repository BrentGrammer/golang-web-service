package main

import (
	"net/http"
	"strconv"
	"fmt"
)

// create methon on our application type that instantiates the routes
func (app *application) route() *http.ServeMux {
		// Create a locally scoped MuxServer
		mux := http.NewServeMux()
		
		// Now create the routes
		
		// call handlefunc off of our serve mux
		mux.HandleFunc("/v1/healthcheck", app.healthcheck) // healthcheck attached as receiver method in handlers.go
		mux.HandleFunc("/v1/books", app.getCreateBooksHandler)
		mux.HandleFunc("/v1/books/", app.getUpdateDeleteBooksHandler)
		// return the mux router
		return mux
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