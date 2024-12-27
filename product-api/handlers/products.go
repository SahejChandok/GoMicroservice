package handlers

import (
	"log"
	"microservicesGo/data"
	"net/http"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if(r.Method == http.MethodGet){
		p.getProducts(rw, r)
		return 
	}

	rw.WriteHeader((http.StatusMethodNotAllowed))
}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	//d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to Marshal Json", http.StatusInternalServerError)
	}
	//rw.Write(d)
}
