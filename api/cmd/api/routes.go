package main

import (
	"net/http"
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