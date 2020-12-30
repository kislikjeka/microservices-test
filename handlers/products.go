package handlers

import (
	"github.com/kislikjeka/microservices/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Product struct {
	logger *log.Logger
}

func NewProducts(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			p.logger.Println("Invalid ID more then one ID")
			http.Error(rw,"Invalid URL",  http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.logger.Println("Invalid ID more then one capture group")
			http.Error(rw,"Invalid URL", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			p.logger.Println("Invalid URl unable to conver num")
			http.Error(rw,"Invalid URL", http.StatusBadRequest)
			return
		}

		p.logger.Println("Got id", id)
		p.updateProduct(id, rw, r)
	}
	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProductsList()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Product) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unmarshal JSON", http.StatusBadRequest)
	}

	data.AddProduct(prod)
	p.logger.Println("Prod: %#v", prod)
}

func (p *Product) updateProduct(id int, rw http.ResponseWriter, r *http.Request){
	p.logger.Println("Handle PUT Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unmarshal JSON", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
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
