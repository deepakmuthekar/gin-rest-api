package main

import (
	"gin-rest-api/controller"
	"log"

	"github.com/gin-gonic/gin"
)

//CustomMiddleware authenticate every request
func CustomMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Inside customMiddleware")
		c.Next()
	}
}

func main() {
	router := gin.Default()

	router.Use(CustomMiddleware())

	api := router.Group("api/v1")
	{
		book := new(controller.BookController)
		api.GET("/books", book.List)
		api.GET("/books/:id", book.Get)
		api.POST("/books", book.Create)
		api.PUT("/books/:id", book.Update)
		api.DELETE("/books/:id", book.Delete)
	}

	router.Run()
}
