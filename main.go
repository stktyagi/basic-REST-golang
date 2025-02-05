package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/books", listBooksHandler)
	r.POST("/books", createBookHandler)
	r.DELETE("/books/:id", deleteBookHandler)

	r.Run()
}

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "title1", Author: "author1"},
	{ID: "2", Title: "title2", Author: "author2"},
}

func listBooksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func createBookHandler(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	books = append(books, book)

	c.JSON(http.StatusCreated, book)
}

func deleteBookHandler(c *gin.Context) {
	id := c.Param("id")

	for i, a := range books {
		if a.ID == id {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
}
