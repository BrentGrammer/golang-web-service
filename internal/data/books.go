package data

import (
	"time"
)

type Book struct {
	ID int64 `json:"id"` // JSON tags allow you to keep idiomatic struct naming conventions, but use JSON conventions when marshalled
	CreatedAt time.Time `json:"-"` // we don't need this in response, so strip it - use directive
	Title string `json:"title"`
	Published int `json:"published",omitempty` // make the field optional by using omit empty
	Pages int `json:"pages",omitempty`
	Genres []string `json:"genres",omitempty`
	Rating float32 `json:"rating",omitempty` // Note: the default val for the data type will be provided: i.e. 0 for float32
	Version int32 `json:"-"` // system data, discard in marshaling for response
}