package main

import (
	"GoJinRestAPI/controllers"
	"GoJinRestAPI/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("books/:id", controllers.DeleteBook)

	if err := r.Run(":8086"); err != nil {
		panic("Error run server!")
	}
}
