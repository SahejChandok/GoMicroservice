package handlers

import (
	"log"
	"microservicesGo/data"
	"net/http"
	"regexp"
	"strconv"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		re := regexp.MustCompile(`/([0-9]+)`)
		g := re.FindAllStringSubmatch(r.URL.Path, -1)
		p.l.Println(g)
		if len(g) != 1 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, _ := strconv.Atoi(idString)
		//p.l.Println("id", id)

		p.updateProducts(id, rw, r)
		return
	}

	//rw.WriteHeader((http.StatusMethodNotAllowed))
}

func (p *Product) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	p.l.Println(err)
	if err != nil {
		http.Error(rw, "Unable to decode", http.StatusBadRequest)
	}
	er := data.UpdateProduct(id, prod)
	if er == data.ErrProductNotFound{
		http.Error(rw, "Product Not found", http.StatusNotFound)
		return
	}
	if er!=nil{
		http.Error(rw, "Product Not found", http.StatusInternalServerError)
		return
	}
}

func (p *Product) addProducts(rw http.ResponseWriter, r *http.Request) {
	//p.l.Println("Handle POST req")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	p.l.Println(err)
	if err != nil {
		http.Error(rw, "Unable to decode", http.StatusBadRequest)
	}
	//p.l.Printf("Product: %#v", prod)
	data.AddProduct(prod)

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
