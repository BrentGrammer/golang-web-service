package main

import (
	"encoding/json"
	"net/http"
)

// helper function for marshaling JSON responses
// attaching to the application struct as a method
func (app *application) writeJSON(w http.ResponseWriter, status int, data any) error {
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