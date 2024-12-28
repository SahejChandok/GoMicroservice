package handlers

import (
	"context"
	"log"
	"microservicesGo/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) UpdateProducts(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	er := data.UpdateProduct(id, prod)

	if er == data.ErrProductNotFound {
		http.Error(rw, "Product Not found", http.StatusNotFound)
		return
	}
	if er != nil {
		http.Error(rw, "Product Not found", http.StatusInternalServerError)
		return
	}
}

func (p *Product) AddProducts(rw http.ResponseWriter, r *http.Request) {
	//p.l.Println("Handle POST req")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	//p.l.Printf("Product: %#v", prod)
	data.AddProduct(prod)

}

func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	//d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to Marshal Json", http.StatusInternalServerError)
		return
	}
	//rw.Write(d)
}

type KeyProduct struct{}

func (p Product) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		p.l.Println(err)
		if err != nil {
			http.Error(rw, "Unable to decode", http.StatusBadRequest)
			return
		}

		// Validate the input 

		err = prod.Validate()
		if err!= nil{
			p.l.Println("error validating product", err)
			http.Error(rw, "Unable to decode", http.StatusBadRequest)
			return 
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
