package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	port := flag.Int("port", 4000, "Number of port to start server")

	flag.Parse()

	r := chi.NewMux()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "teste")
	})

	fmt.Printf("Start server on port: %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
