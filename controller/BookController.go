package controller

import (
	"gin-rest-api/api"
	"gin-rest-api/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//BookController  BookController
type BookController struct{}

//List : Return list of Books
func (bc *BookController) List(c *gin.Context) {
	repository := new(repository.BookRepository)
	var books = repository.List()
	c.JSON(200, books)
}

//Get : Return book by Id
func (bc *BookController) Get(c *gin.Context) {
	id := c.Param("id")
	repository := new(repository.BookRepository)
	book := repository.Get(id)
	c.JSON(http.StatusOK, book)

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
	repository := new(repository.BookRepository)

	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		bookID := repository.Create(book)
		//TODO : send location Header
		c.JSON(http.StatusOK, bookID)
	}
}
