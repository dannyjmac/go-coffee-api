// CREATE A BASIC WEB SERVER

// To Start
// go run main.go

// Runing/Testing
// GET all items = curl localhost:9090 | jq
// POST an item = curl localhost:9090 -d ' {"name": "Drink"}'

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dannyjmac/go-micro-3/handlers"
)

func main() {

	l := log.New(os.Stdout, "products-api", log.LstdFlags)
	ph := handlers.NewProducts(l)

	// This is a servemux, basically a router like in node.js
	sm := http.NewServeMux()
	// Like nodejs routes, when a request comes into the server, calls the products handler
	sm.Handle("/", ph)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// A goroutine - how go handles concurrency
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	s.ListenAndServe()

	// A shutdown function that waits until the requests have been completed before shutting down gracefully
	// So that there are not clients trying to fetch data as you shut the server
	// In this case take 30 seconds to shutdown
	tc, _ := context.WithTimeout(context.Background(), time.Second*30)
	s.Shutdown(tc)

}
