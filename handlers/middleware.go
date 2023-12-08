package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	token string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{
		token: os.Getenv("AUTH_TOKEN"),
	}

	if config.token == "" {
		log.Fatal("No AUTH_TOKEN found in .env file")
	}

	return config
}

func (c *Config) authenticateAndMethodMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			log.Printf("Method Not Allowed - Origin: %s, Method: %s", r.Header.Get("Origin"), r.Method)
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		token := r.Header.Get("Authorization")

		if token != c.token {
			log.Printf("Unauthorized Access Attempt - Origin: %s", r.Header.Get("Origin"))
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		handler(w, r)
	})
}

func trustedDomainMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		originAllowed := false

		// Add the trusted origins
		allowedOrigins := []string{
			"http://localhost:3000",
			"http://localhost:8000",
			"http://front-end",
		}

		// Check if the request's origin matches any allowed origin
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				originAllowed = true
				break
			}
		}

		// Allow requests without a specified Origin (local access)
		if origin == "" {
			originAllowed = true
		}

		if !originAllowed {
			log.Printf("Forbidden - Unauthorized Origin: %s", r.Header.Get("Origin"))
			http.Error(w, "Forbidden - Unauthorized", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
