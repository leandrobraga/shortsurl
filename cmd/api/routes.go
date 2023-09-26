package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	r := chi.NewMux()
	r.Use(middleware.Logger)

	r.Get("/{code}", app.redirectUrlHandler)
	r.Post("/", app.createshortUrlHandler)

	return r
}
