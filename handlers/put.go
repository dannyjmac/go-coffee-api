package handlers

import (
	"net/http"
	"strconv"

	"github.com/dannyjmac/go-micro-3/data"
	"github.com/gorilla/mux"
)

// swagger:route PUT /products{id} products updateProduct
// Updates a product
// responses:
//	201: noContent
// 	422: errorValidation
//	501: errorResponse

// Update updates a product in the database
func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
	}

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
	}
}
