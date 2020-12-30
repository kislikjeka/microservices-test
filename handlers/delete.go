package handlers

import (
	"github.com/gorilla/mux"
	"github.com/kislikjeka/microservices/data"
	"net/http"
	"strconv"
)

func (p Product) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Invalid convertinc id to int", http.StatusBadRequest)
	}
	p.logger.Println("Handle DELETE Product")

	err = data.DeleteProduct(id)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
