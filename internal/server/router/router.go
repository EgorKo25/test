package router

import (
	"test/internal/server/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(handler *handler.Handler) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.HandleFunc("/solve", handler.LineS)

	return r
}
