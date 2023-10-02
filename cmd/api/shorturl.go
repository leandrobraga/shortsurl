package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/leandrobraga/shortsurl/internal/codegenerator"
	"github.com/leandrobraga/shortsurl/internal/data"
)

func (app *application) createshortUrlHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Url string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url := &data.ShortUrl{
		Url:  input.Url,
		Code: codegenerator.CodeGenerator(6),
	}
	ok, errors := url.IsValid()
	if !ok {
		data, err := json.Marshal(errors)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}

	shorUrlModel := data.ShortUrlModel{
		DB: app.db,
	}
	err := shorUrlModel.Insert(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{"code":"%s"}`, url.Code)))
}

func (app *application) redirectUrlHandler(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	_ = code
	http.Redirect(w, r, "https://google.com.br", http.StatusTemporaryRedirect)
}
