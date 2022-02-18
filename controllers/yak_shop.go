package controllers

import (
	"github.com/gin-gonic/gin"
	"lulu/go-yakshop/models"
	"net/http"
)

type CreateYakRequest struct {
	Name string `json:"name" binding:"required"`
	Age  uint   `json:"age" binding:"required"`
}

func GetYaks(c *gin.Context) {
	var yaks []models.Yak
	models.DB.Find(&yaks)

	c.JSON(http.StatusOK, yaks)
}

func CreateYak(c *gin.Context) {
	var newYak CreateYakRequest
	if err := c.ShouldBindJSON(&newYak); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Yak{Name: newYak.Name, Age: newYak.Age}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, newYak)
}

func GetYakByName(c *gin.Context) {
	name := c.Param("name")

	var yak models.Yak
	if err := models.DB.First(&yak, "name = ?", name).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "yak not found"})
		return
	}

	c.JSON(http.StatusOK, yak)
}
