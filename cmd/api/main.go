package main

import (
	"flag"
	"fmt"
	"net/http"
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
func main() {
	// use our config struct to set details of our app
	var cfg config
	// use the flag package to modify config details
	// pass in mem addr of the config entries so we modify them, name of property, value and description/help
	flag.IntVar(&cfg.port, "port", 4000, "API Server Port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stage|prod)")
	// need to call Parse() for flag package to modify the config
	flag.Parse()

	
	//instantiate a logger - write to stdout and set the date and time
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	
	app := &application{
		config: cfg, // the config set with the flag package above
		logger: logger,
	}
	
	// create var for the port using our config setting
	addr := fmt.Sprintf(":%d", cfg.port)

	// Create a locally scoped MuxServer
	mux := http.NewServeMux()
	// call handlefunc off of our serve mux
	mux.HandleFunc("/v1/healthcheck", app.healthcheck) // healthcheck attached as receiver method in handlers.go

	err := http.ListenAndServe(addr, mux) // Note: passing nil uses a default serve mux so you don't have to create one
	// when using default ServeMux it is a global variable so it's possible to inject handlers by just using http.HandleFunc anywhere without specifying a server
	// To get around this, we can create our locally scoped serve mux ourselves and pass it in as shown here.

	logger.Printf("Starting %s server on port %s", cfg.env, addr)
	if err != nil {
		fmt.Println(err)
	}
}