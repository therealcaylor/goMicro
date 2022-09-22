package handlers

import (
	"gomicro/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l: l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	//! old way
	// data, err := json.Marshal(lp)
	// if err != nil {
	// 	http.Error(w, "not able to marshal JSON", http.StatusInternalServerError)
	// }
	// w.Write(data)
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "not able to marshal JSON", http.StatusInternalServerError)
	}
}
