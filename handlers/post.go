package handlers

import (
	"github.com/kislikjeka/microservices/data"
	"net/http"
)

//AddProduct add product to the data store
func (p *Product) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)
	p.logger.Println("Prod: %#v", prod)
}
