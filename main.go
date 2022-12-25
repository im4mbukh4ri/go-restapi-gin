package main

import (
	"ecommerce/controllers"
	"ecommerce/models"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	models.ConnectDatabase()

	route.GET("/api/products", controllers.Index)
	route.GET("/api/product/:id", controllers.Show)
	route.POST("/api/product", controllers.Store)
	route.PUT("/api/product/:id", controllers.Update)
	route.DELETE("/api/product", controllers.Delete)

	route.Run()
}
