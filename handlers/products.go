package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/dannyjmac/go-micro-3/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// Check if its a get request
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
	}

	if r.Method == http.MethodPut {
		// expect the id in the URI

		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(rw, "Must provide an ID", http.StatusBadRequest)
		}

		if len(g[0]) != 2 {
			http.Error(rw, "More than one caputure group", http.StatusBadRequest)
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(rw, "Invalid ID", http.StatusBadRequest)
			return
		}

		p.updateProducts(id, rw, r)
		return
	}

	// If not a get request, return method is not allowed
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	// Take data in the post and convert it to our struct - a JSON encoder
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to interpret your JSON, Jason", http.StatusBadRequest)
	}

	data.AddProduct(prod)

}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	// Take data in the post and convert it to our struct - a JSON encoder
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to interpret your JSON, Jason", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
	}
}
