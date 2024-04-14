package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
	"log"
	"os"
)

const version = "1.0.0"

type config struct{
	port int 
	env string
}

type application struct {
	config config
	logger *log.Logger
}
// use our config struct to set details of our app
var cfg config

func main() {
	// use the flag package to modify config details
	// pass in mem addr of the config entries so we modify them, name of property, value and description/help
	flag.IntVar(&cfg.port, "port", 4000, "API Server Port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stage|prod)")
	// need to call Parse() for flag package to modify the config
	flag.Parse()

	
	//instantiate a logger - write to stdout, specify a prefix string, and set the date and time
	// The | (bitwise OR) operator is used to combine these flags into a single value. This means that the logger will prepend each log message with the date and time.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime) // Ldate and Ltime is local date and local time
	
	app := &application{
		config: cfg, // the config set with the flag package above
		logger: logger,
	}
	
	// create var for the port using our config setting
	addr := fmt.Sprintf(":%d", cfg.port)

	// create server type that instantiates our separated routes in routes.go and sets other params
	srv := &http.Server{
		Addr: addr,
		Handler: app.route(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on port %s", cfg.env, addr)
	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}