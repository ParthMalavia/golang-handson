package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"example/go-bookstore/pkg/models"
	"example/go-bookstore/pkg/utils"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	allBooks := models.GetAllBooks()
	jsonAllBooks, _ := json.Marshal(allBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonAllBooks)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing BookID:", err)
	}
	b, _ := models.GetBookByID(ID)
	jsonBook, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBook)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	b := newBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID := mux.Vars(r)["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing BookID:", err)
	}
	b := models.DeleteBook(ID)
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}

// TODO: Try with different logic
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updatedBook = &models.Book{}
	utils.ParseBody(r, updatedBook)
	bookID := mux.Vars(r)["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing BookID:", err)
	}
	fetchedBook, db := models.GetBookByID(ID)
	if updatedBook.Name != "" {
		fetchedBook.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		fetchedBook.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		fetchedBook.Publication = updatedBook.Publication
	}
	db.Save(&fetchedBook)
	res, _ := json.Marshal(fetchedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
