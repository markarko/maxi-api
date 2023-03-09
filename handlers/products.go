package handlers

import (
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
		p.l.Println(r.Header)
		if r.Header.Get("X-Api-Key") == "" || !ValidApiKey(r.Header.Get("X-Api-Key")) {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		products := maxi.GetAllProducts("eggs")
		if products == nil {
			http.Error(rw, "Unable to get products", http.StatusInternalServerError)
		}

		err := products.ToJSON(rw)

		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusInternalServerError)
		}
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}
