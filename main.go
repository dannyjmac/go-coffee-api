// CREATE A BASIC WEB SERVER

// To Start
// go run main.go

// Runing/Testing
// GET all items = curl localhost:9090 | jq
// POST an item = curl localhost:9090 -d ' {"name": "Drink"}'
// PUT (update an item) = curl localhost:9090/1 -XPUT -d '{"name": "fanta"}'

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dannyjmac/go-micro-3/handlers"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "products-api", log.LstdFlags)
	ph := handlers.NewProducts(l)

	// Gorilla's mux server
	sm := mux.NewRouter()

	// Gives a router specfivally for a GET route
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

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
