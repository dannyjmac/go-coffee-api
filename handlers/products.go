package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dannyjmac/go-micro-3/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// Previously private methods (getProducts etc) have now been made public
// As they are no longer being called buy the ServeHTTP method we had. Gorilla
// Is handling all of the serving
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	// Take data in the post and convert it to our struct - a JSON encoder
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to interpret your JSON, Jason", http.StatusBadRequest)
	}

	data.AddProduct(prod)

}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	// The id of the product to update now comes from Gorilla mux
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
	}


	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
	}
}

type KeyProduct struct

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandleFunc(rw *http.ResponseWriter, r *http.Request) {
		// Take data in the post and convert it to our struct - a JSON encoder
		prod := &data.Product{}

		err = prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to interpret your JSON, Jason", http.StatusBadRequest)
			return
		}

		ctx := r.Context().WithValue(KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	}
}
