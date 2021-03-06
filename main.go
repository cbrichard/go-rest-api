package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	//	"math/rand"
	"net/http"
	//	"strconv"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
}

// Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {
}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data - @todo - implement DB
	books = append(books, Book{ID: "1", ISBN: "448743", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", ISBN: "448877", Title: "Book Two", Author: &Author{Firstname: "Jane", Lastname: "Doe"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books/", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Run Server with logging
	log.Fatal(http.ListenAndServe(":8000", r))
}
