package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"strconv"
	// "io/ioutil"

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

// get books list or create a new book
func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	// accepts GET and POST requests
	if r.Method == http.MethodGet {
		books := []data.Book{
			{
				ID: 1,
				CreatedAt: time.Now(),
				Title: "The Darkening of Trstram",
				Published: 1998,
				Pages: 300,
				Genres: []string{"Fiction","Thriller"},
				Rating: 4.5,
				Version: 1,
			},
			{
				ID: 2,
				CreatedAt: time.Now(),
				Title: "The legacy of Deckard",
				Published: 2007,
				Pages: 232,
				Genres: []string{"Fiction","Adventure"},
				Rating: 4.9,
				Version: 1,
			},
		}
        // use helper to marshal json and return json response
		// this also uses the inline scoped variable in if statement (var; condition to check)
		if err := app.writeJSON(w, http.StatusOK, envelope{"books": books}); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	
	if r.Method == http.MethodPost {
		// contains details of a book structure that we expect from a request that we will unmarshal into a go object
		var input struct {
			Title string `json:"title"`
			Published int `json:"published"`
			Pages int `json:"pages"`
			Genres []string `json:"genres"`
			Rating float64 `json:"rating"`
		}

		// old way of getting body - left for reference
		// get the body of the request
		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		// 	return 
		// }
		// // we want to convert the body of the request to a go struct, pass in mem addr of the input struct to mutate it
		// err = json.Unmarshal(body, &input)

		err := app.readJSON(w, r, &input)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "%v\n", input)
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

	book := data.Book{
		ID: idInt,
		CreatedAt: time.Now(),
		Title: "Echoes in the Darkness",
		Published: 2019,
		Pages: 300,
		Genres: []string{"Fiction","Thriller"},
		Rating: 4.5,
		Version: 1,
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"book": book}); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) updateBook(w http.ResponseWriter, r *http.Request) {
    // get the id of the book from the url string
	id := r.URL.Path[len("/v1/books/"):] // grab the part of the path after the route
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		// if we can't parse to integer we can't continue
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	
	// Use pointers so that we can modify the existing struct in place instead of having to create a new one.
	var input struct {
		Title *string `json:"title"` // in this case json tags tell us what we look for in the request json
		Published *int `json:"published"`
		Pages *int `json:"pages"`
		Genres []string `json:"genres"`
		Rating *float32 `json:"rating"`
	}

	// this is just a mock book that acts as the existing record
	book := data.Book{
		ID: idInt,
		CreatedAt: time.Now(),
		Title: "Echoes in the Darkness",
		Published: 2019,
		Pages: 300,
		Genres: []string{"Fiction","Thriller"},
		Rating: 4.5,
		Version: 1,
	}

	// Left for reference (old way of reading json request)
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// 	return 
	// }

	// unmarshal the info from the request into the input struct
	// err = json.Unmarshal(body, &input) // load data into the input struct
	// if err != nil{
	// 	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// }

	err = app.readJSON(w, r, &input) // pass a pointer (mem addr) of input
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

    // check the input values and update the record if they are not nil
	if input.Title != nil{
		book.Title = *input.Title 
	}
	if input.Published != nil{
		book.Published = *input.Published
	}
	if input.Pages != nil {
		book.Pages = *input.Pages
	}
	if len(input.Genres) > 0{
		book.Genres = input.Genres
	}
	if input.Rating != nil{
		book.Rating = *input.Rating
	}

	fmt.Fprintf(w, "%v\n", book)
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

