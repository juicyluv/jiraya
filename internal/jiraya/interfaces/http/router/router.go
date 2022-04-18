package router

import "github.com/go-chi/chi"

func New() *chi.Mux {
	r := chi.NewRouter()

	return r
}
