package main

import (
	"encoding/json"
	"net/http"
	"io"
	"errors"
)

// wrap the response data in an object with a key for the data: i.e. {  book: bookData }
type envelope map[string]any

// helper function for marshaling JSON responses
// attaching to the application struct as a method
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status) // set the http status of the response
	w.Write(js)

	return nil
}

// helper to use when we marshal JSON into a go object/struct (i.e. for reading request bodies)
// dst can be a pointer to some request input so we can mutate it here
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
    // protect endpoint with a maximum bytes allowed in the request
	maxBytes := 1_048_576

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// *** need to switch to using decoder instead of Unmarshal to set the max bytes and disallow unknown fields and throw errors
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields() // can't pass in fields that are not in the interface struct for the request

	// decode the body into the dst pointer passed in
	if err := dec.Decode(dst); err != nil {
		return err
	}
    
	// This is done to ensure that there are no additional JSON objects or data present after the primary object. If decoding the empty struct does not result in an io.EOF error, it means there is extra data in the request body, which violates the expectation of a single JSON object.
	err := dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON object")
	}
	
	return nil
}