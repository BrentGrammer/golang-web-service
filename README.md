# Web Service created with GoLang (WIP)

- Based on the [Pluralsight exercise](https://app.pluralsight.com/library/courses/go-building-web-services-applications/table-of-contents) by Josh Duffney
- The code is heavily commented for reference and this is used as a learning and experimentation project.
- This is a learning project and is in a WIP state. It is meant for experimenting and for reference.

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

### internal folder

- "internal" means that any package here cannot be imported from outside this project (the code/files are not relevant to any outside project but our own - they are internal to the project)
- `internal/data`: Use this folder to hold data types and models

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

## Database Setup

- Start the Docker container with Postgres
- `$ psql -h localhost -p 5432 -U postgres`
- `CREATE DATABASE readinglist;`
- `CREATE ROLE readinglist WITH LOGIN PASSWORD 'pa55w0rd';`
- `\c readinglist` change to the new db
- Create tables

```sql
CREATE TABLE IF NOT EXISTS books (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    published integer NOT NULL,
    pages integer NOT NULL,
    genres text[] NOT NULL,
    rating real NOT NULL,
    version integer NOT NULL DEFAULT 1
);
```

- grant permissions

```sql
GRANT SELECT, INSERT, UPDATE, DELETE ON books TO readinglist;
```

- for using primary key with bigserial we need additional permissions

```sql
GRANT USAGE, SELECT ON SEQUENCE books_id_seq TO readinglist;
```
