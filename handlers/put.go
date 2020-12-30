package handlers

import (
	"github.com/gorilla/mux"
	"github.com/kislikjeka/microservices/data"
	"net/http"
	"strconv"
)

func (p *Product) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Invalid convertinc id to int", http.StatusBadRequest)
	}
	p.logger.Println("Handle PUT Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
	p.logger.Println("Prod: %#v", prod)
}
