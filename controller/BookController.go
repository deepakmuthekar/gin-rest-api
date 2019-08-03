package controller

import (
	"gin-rest-api/api"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//BookController  BookController
type BookController struct{}

//List : Return list of Books
func (bc *BookController) List(c *gin.Context) {
	var books = []api.Book{
		api.Book{Title: "Scala In Action", Author: "Martin", Pages: 200},
		api.Book{Title: "Scala In Action", Author: "Martin", Pages: 200},
	}
	c.JSON(200, books)
}

//Get : Return book by Id
func (bc *BookController) Get(c *gin.Context) {
	id := c.Param("id")
	log.Println("Inside get, Getting book with id " + id)
}

//Update : Update book by Id
func (bc *BookController) Update(c *gin.Context) {
	log.Println("Inside update")
}

//Delete : Delete book by Id
func (bc *BookController) Delete(c *gin.Context) {
	log.Println("Inside delete")
}

//Create : Creates bew  Book.
func (bc *BookController) Create(c *gin.Context) {
	var book api.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, book)
	}
}
