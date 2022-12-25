package controllers

import (
	"../../models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOk, gin.H{"products": products})
}

func Show(c *gin.Context) {

}

func Store(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
