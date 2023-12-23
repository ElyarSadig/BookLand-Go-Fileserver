package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Routes() http.Handler {

	router := mux.NewRouter()

	// CORS middleware for all routes
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081", "http://localhost:3000", "*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	router.Use(corsMiddleware.Handler)

	router.PathPrefix("/identities/").Handler(http.StripPrefix("/identities/", http.FileServer(http.Dir("uploads/identities"))))

	router.PathPrefix("/publications/").Handler(http.StripPrefix("/publications/", http.FileServer(http.Dir("uploads/publications"))))

	router.PathPrefix("/book-covers/").Handler(http.StripPrefix("/book-covers/", http.FileServer(http.Dir("uploads/book_covers"))))

	router.PathPrefix("/books/").Handler(http.StripPrefix("/books/", http.FileServer(http.Dir("uploads/books"))))

	config := NewConfig()

	router.HandleFunc("/upload/identities", config.authenticateAndMethodMiddleware(uploadIdentityHandler))
	router.HandleFunc("/upload/publications", config.authenticateAndMethodMiddleware(uploadPublicationHandler))
	router.HandleFunc("/upload/book-covers", config.authenticateAndMethodMiddleware(uploadBookCoverHandler))
	router.HandleFunc("/upload/books", config.authenticateAndMethodMiddleware(uploadBookHandler))

	return router
}
