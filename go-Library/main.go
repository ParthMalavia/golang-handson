package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	qty    int
}

// var books = []book{
// 	{ID: "1", Title: "IN Search of Lost Time", Author: "Marshal Proust", qty: 2},
// 	{ID: "2", Title: "Angels And Demons", Author: "Dan Brown", qty: 7},
// 	{ID: "3", Title: "Study In Scarlet", Author: "Arthur Conan Doyle", qty: 6},
// }

var books = map[string]*book{
	"1": {ID: "1", Title: "IN Search of Lost Time", Author: "Marshal Proust", qty: 2},
	"2": {ID: "2", Title: "Angels And Demons", Author: "Dan Brown", qty: 7},
	"3": {ID: "3", Title: "Study In Scarlet", Author: "Arthur Conan Doyle", qty: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookById(c *gin.Context) {
	id := c.Param("id")
	book, available := books[id]

	if !available {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func validateBook(c *gin.Context) *book {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Missing query param `id`."})
		return nil
	}

	book, available := books[id]
	if !available {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not found."})
		return nil
	}

	return book
}

func checkoutBook(c *gin.Context) {
	book := validateBook(c)

	if book.qty <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book currently not in stock."})
		return
	}

	book.qty--
	c.IndentedJSON(http.StatusOK, gin.H{"book": book, "Remaining in stock": book.qty})
}

func returnBook(c *gin.Context) {
	book := validateBook(c)
	book.qty++
	c.IndentedJSON(http.StatusOK, *book)
}

func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		fmt.Println(err)
		return
	}
	// books = append(books, newBook)
	books[newBook.ID] = &newBook
	c.IndentedJSON(http.StatusCreated, gin.H{"ID": newBook.ID, "Title": newBook.Title})
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/book/:id", getBookById)
	router.POST("/book", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run(":8080")
}
