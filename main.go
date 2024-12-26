package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello World")
		d, err := io.ReadAll(r.Body)
		if err!= nil{
			http.Error(w, "oops", http.StatusBadRequest)
			return 
		}
		fmt.Fprintf(w, "Hello %s", d)
	})
	http.HandleFunc("/GoodBye", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Goodbye")
	})

	http.ListenAndServe(":9090", nil)
}
