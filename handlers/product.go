package handlers

import (
	"log"
	"net/http"

	"github.com/thisisrahmat/microservices_in_go/product-api/data"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// interface
// so within it you handle different requests and return different responsewriter
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle the request for a list of products
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// if r.Method == http.MethodPost {
	// 	p.addProducts(rw, r)
	// 	return
	// }

	// catch all
	// if the GET method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from the data store
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	//get list of products from abstracted GetProducts method
	//lp stands for list of products
	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
// 	p.l.Println("Handle POST Products")

// 	//create new product object

// 	// why do we need & and .

// 	prod := &data.Product{}

// 	err := prod.ToJSON(r.Body)

// 	if err != nil {
// 		http.Error(rw, "Unable to unmarshal JSON", http.StatusInternalServerError)
// 	}

// 	p.l.Printf("Prod: %#v", prod)

// }
