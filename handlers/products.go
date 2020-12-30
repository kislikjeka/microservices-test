// Package classification of Product API
//
// Documentation Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
//  - application/json
//
// Produces:
//  - application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"github.com/kislikjeka/microservices/data"
	"log"
	"net/http"
)

// Product is  http.Handler
type Product struct {
	logger *log.Logger
}

func NewProducts(l *log.Logger) *Product {
	return &Product{l}
}

type KeyProduct struct{}

func (p Product) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.logger.Println("[ERROR] deserialazing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			p.logger.Println("[ERROR] validating product", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
