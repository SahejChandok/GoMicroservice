package main

import (
	"context"
	"log"
	"microservicesGo/handlers"
	"net/http"
	"os"
	"time"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// hh := handlers.NewHello(l)
	// gh := handlers.NewGoodbye(l)
	ph := handlers.NewProduct(l)
	sm := http.NewServeMux()


	// sm.Handle("/", hh)
	// sm.Handle("/bye", gh)
	sm.Handle("/", ph)


	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout: 120 * time.Second,
	}

	s.ListenAndServe()
 
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
