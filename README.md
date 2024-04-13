# Web Service created with GoLang

- Based on the [Pluralsight exercise](https://app.pluralsight.com/library/courses/go-building-web-services-applications/table-of-contents) by Josh Duffney
- The code is heavily commented for reference and this is used as a learning and experimentation project.

## Starting the Application

- Requirements: Docker (Docker Desktop)
- `docker compose up --build`

## Project Structure

### cmd folder

- Files in this folder should be compiled to a binary
- Can be moved into a bin directory later on

### Routes
- Place in `cmd/api/routes.go` to separate routes
- create mux server in here for the routes
- Move calls to HandlFunc() into this file

### cmd/api/handlers.go
- Handlers are decoupled into the cmd/api/handlers.go file

## Returning Errors

### Constraining what methods are allowed for an Endpoint:
```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    if (r.Method != http.MethodGet) {
        http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
        return // make sure to return and exit the handler!
    }
}
```