package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/thisisrahmat/microservices_in_go/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// hh := handlers.NewHello(l)

	gh := handlers.NewGoodbye(l)

	//create the handlers
	ph := handlers.NewProducts(l)

	//http.HandleFunc - covnerts function into a handler type
	//and then registerign it into the defaultServerMux
	//default sevrerMux contains logic to know which handler to call based on the path
	//serverMux - is an object/type

	//internally what is happening is that when a request comes into your sevrer,
	// the server has a default handler
	//the fefauly handler is the http.serveMux

	//special logic

	sm := http.NewServeMux()
	sm.Handle("/", ph)
	sm.Handle("/goodbye", gh)

	// starts the http server, defaultsevrermux
	//takes two parameters first is binding address then secodn is the http handler
	//if you don;t sepcfiy the handler it uses the default serveMux
	// http.ListenAndServe(":9090", sm)

	//manually create a HTTP server and then feed that into
	//listen and serve so you can control and fine tune things

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	s.ListenAndServe()

	//shutdown server gracefully using Shutdown function
	//the shutdown when called no longer accepts any new requests but waits until it's
	// all the existing requests are finished before shutting down

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
