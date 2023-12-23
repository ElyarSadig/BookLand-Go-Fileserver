package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func uploadIdentityHandler(w http.ResponseWriter, r *http.Request) {
	uploadHandler(w, r, "uploads/identities")
}

func uploadPublicationHandler(w http.ResponseWriter, r *http.Request) {
	uploadHandler(w, r, "uploads/publications")
}

func uploadBookCoverHandler(w http.ResponseWriter, r *http.Request) {
	uploadHandler(w, r, "uploads/book_covers")
}

func uploadBookHandler(w http.ResponseWriter, r *http.Request) {
	uploadHandler(w, r, "uploads/books")
}

func uploadHandler(w http.ResponseWriter, r *http.Request, uploadDir string) {
	r.ParseMultipartForm(10 << 20) // 10 MB

	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println("Error retrieving the file:", err)
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Check the file extension and handle accordingly
	extension := filepath.Ext(handler.Filename)
	switch extension {
	case ".pdf", ".png", ".jpg", ".jpeg":
		// Handle supported file types
	default:
		log.Println("Unsupported file type: ", extension)
		http.Error(w, "Unsupported file type", http.StatusBadRequest)
		return
	}

	// Create the directory if it doesn't exist
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, os.ModePerm)
	}

	// Create a new file in the server's filesystem to store the uploaded file
	filePath := filepath.Join(uploadDir, handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		log.Println("Error creating the file:", err)
		http.Error(w, "Error creating the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the file to the destination
	_, err = io.Copy(dst, file)
	if err != nil {
		log.Println("Error copying the file:", err)
		http.Error(w, "Error copying the file", http.StatusInternalServerError)
		return
	}

	log.Printf("File uploaded - Origin: %s, File: %s, Path: %s", r.Header.Get("Origin"), handler.Filename, filePath)
	// Respond to the client with a success message
	fmt.Fprintf(w, "File uploaded successfully to: %s", filePath)
}
