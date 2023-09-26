package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) createshortUrlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("here the code in json"))
}

func (app *application) redirectUrlHandler(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	_ = code
	http.Redirect(w, r, "https://google.com.br", http.StatusTemporaryRedirect)
}
