package handlers

import (
	"github.com/kislikjeka/microservices/data"
	"net/http"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//	200: productsResponse

//GetProducts return products from the data store
func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle GET Products")

	// Fetch the products from data store
	lp := data.GetProductsList()

	// serialize list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
