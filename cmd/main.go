package main

import (
	"fmt"
	"net/http"

	"github.com/elyarsadig/file_server/file_operations"
	"github.com/elyarsadig/file_server/handlers"
)

func main() {
	// Create a directory for uploads if it doesn't exist
	directories := []string{
		"uploads",
		"uploads/identities",
		"uploads/publications",
		"uploads/book_covers",
		"uploads/books",
	}

	// Create directories if they don't exist
	fileoperations.CreateFilePaths(directories)

	handlers.InitializeFileServers()

	handlers.SetupUploadHandlers()

	// Start the server on port 8080
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
