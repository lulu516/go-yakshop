package main

import (
	"github.com/gin-gonic/gin"
	"lulu/go-yakshop/controllers"
	"lulu/go-yakshop/models"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	router.GET("/yaks", controllers.GetYaks)
	router.GET("/yaks/:name", controllers.GetYakByName)
	router.POST("/yak", controllers.CreateYak)

	router.Run("localhost:8080")
}
