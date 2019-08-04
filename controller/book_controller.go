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
// @Summary List all books
// @Description List all books
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {object} api.Book
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /books/ [get]
func (bc *BookController) List(c *gin.Context) {
	repository := new(repository.BookRepository)
	var books = repository.List()
	c.JSON(200, books)
}

//Get : Return book by Id
// @Summary Return Book by Id
// @Description Return Book by Id
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} api.Book
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /books/{id} [get]
func (bc *BookController) Get(c *gin.Context) {
	id := c.Param("id")
	repository := new(repository.BookRepository)
	book := repository.Get(id)
	c.JSON(http.StatusOK, book)
}

//Update : Update book by Id
// @Summary Update Book details
// @Description Update Book details
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body api.Book true "Update Book"
// @Success 200 {object} api.Book
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /books/ [put]
func (bc *BookController) Update(c *gin.Context) {
	log.Println("Inside Update")
	var book api.Book
	repository := new(repository.BookRepository)

	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		_, err := repository.Update(book)

		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Book record uodated succesfully !!!",
			})
		}
	}

}

//Delete : Delete book by Id
// @Summary Delete Book record by Id
// @Description Delete Book record by Id
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} api.Book
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /books/{id} [delete]
func (bc *BookController) Delete(c *gin.Context) {
	log.Println("Inside delete")
	id := c.Param("id")
	repository := new(repository.BookRepository)

	_, err := repository.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Book with ID '" + id + "' deleted succesfully !!!",
		})
	}

}

//Create : Creates bew  Book.
// @Summary Create new Book record in Library
// @Description Create new Book record in Library
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body api.Book true "Add Book"
// @Success 200 {object} api.Book
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /books/ [post]
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
