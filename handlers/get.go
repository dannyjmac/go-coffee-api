package handlers

import (
	"net/http"

	"github.com/dannyjmac/go-micro-3/data"
)

// swagger:route GET /products products listProducts
// Returns a list of Products
// responses:
//	200: productsResponse

// ListProducts returns the products from the data store
func (p *Products) ListProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

}
