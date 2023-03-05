package handlers

import (
	//"encoding/json"
	"log"
	"net/http"

	"github.com/markarko/maxi-api/maxi"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		//prod := &data.Product{}
		products := maxi.GetAllProducts("eggs")
		if products == nil {
			http.Error(rw, "Unable to get products", http.StatusInternalServerError)
		}

		err := products.ToJSON(rw)

		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusInternalServerError)
		}
	}
	//rw.WriteHeader(http.StatusMethodNotAllowed)
}
