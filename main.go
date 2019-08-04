package main

import (
	"gin-rest-api/controller"
	"log"

	"github.com/gin-contrib/cors"

	_ "gin-rest-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

//CustomMiddleware authenticate every request
func CustomMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Inside customMiddleware")
		c.Next()
	}
}

// @title Library  API
// @version 1.0
// @license.name Apache 2.0
// @description Library API Server.
// @host localhost:8080
// @contact.name Deepak Muthekar
// @BasePath /api/v1
func main() {

	router := gin.Default()

	router.Use(CustomMiddleware())
	router.Use(cors.Default())

	api := router.Group("api/v1")
	{
		book := new(controller.BookController)
		api.GET("/books", book.List)
		api.GET("/books/:id", book.Get)
		api.POST("/books", book.Create)
		api.PUT("/books", book.Update)
		api.DELETE("/books/:id", book.Delete)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}
