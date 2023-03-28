package main

import (
	"chap2-project/controller"
	"chap2-project/database"
	"chap2-project/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	database.StartDB()
	db := database.GetDB()

	bookRepository := repository.NewBookRepository(db)
	bookController := controller.NewBookController(*bookRepository)

	g := gin.Default()
	g.GET("/books", bookController.GetBooks)
	g.GET("/books/:id", bookController.GetBookById)
	g.POST("/books", bookController.AddBook)
	g.PUT("/books/:id", bookController.UpdateBook)
	g.DELETE("/books/:id", bookController.DeleteBook)

	g.Run(":8080")
}
