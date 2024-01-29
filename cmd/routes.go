package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	mux.Get("/", app.Home)

	// Books routes
	mux.Route("/book", func(mux chi.Router) {
		mux.Get("/", app.Book)
		mux.Post("/all", app.AllBooks)
		mux.Put("/edit/{id}", app.EditBook)
		mux.Post("/save", app.SaveBook)
		mux.Delete("/delete/{id}", app.DeleteBook)
		mux.Post("/{id}", app.GetBook)
	})
	mux.Route("/sale", func(mux chi.Router) {
		// mux.Get("/", app.Sale)
		// mux.Post("/all", app.AllSales)
		// mux.Put("/edit/{id}", app.EditSale)
		mux.Post("/save", app.SaveSale)
		// mux.Delete("/delete/{id}", app.DeleteSale)
		// mux.Post("/{id}", app.GetSale)
	})

	// Donate routes
	mux.Get("/donate", app.Donate)

	fileServer := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/*", http.StripPrefix("/assets", fileServer))

	return mux
}
