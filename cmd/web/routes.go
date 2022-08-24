package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jacstn/golang-buy-btc-from-me/internal/handlers"
)

func routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(LoadSession)
	mux.Use(NoSurf)
	mux.Get("/", handlers.Home)
	mux.Post("/create-order", handlers.CreateOrder)
	mux.Post("/update-order-status", handlers.UpdateOrderStatus)
	mux.Get("/get-btc-price", handlers.GetBTCPrice)
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
