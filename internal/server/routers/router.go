package routers

import (
	"test/internal/server/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New(hdr *handlers.Handler) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.HandleFunc("/solve", hdr.LineS)

	return r
}
