package handlers

import (
	"net/http"

	"github.com/dannyjmac/go-micro-3/data"
)

// swagger:route POST /products products createProduct
// Creates a product
// responses:
//	200: productResponse
// 	422: errorValidation
//	501: errorResponse

// Create returns the products from the data store
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)

}
