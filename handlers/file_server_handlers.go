package handlers

import "net/http"


func InitializeFileServers() {
	
	http.Handle("/identities/", trustedDomainMiddleware(http.StripPrefix("/identities/", http.FileServer(http.Dir("uploads/identities")))))

	http.Handle("/publications/", trustedDomainMiddleware(http.StripPrefix("/publications/", http.FileServer(http.Dir("uploads/publications")))))

	http.Handle("/book-covers/", trustedDomainMiddleware(http.StripPrefix("/book-covers/", http.FileServer(http.Dir("uploads/book_covers")))))

	http.Handle("/books/", trustedDomainMiddleware(http.StripPrefix("/books/", http.FileServer(http.Dir("uploads/books")))))

}
