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
	r.Run(":8086")
}
