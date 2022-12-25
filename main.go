package main

import (
	"github.com/im4mbukh4ri/go-restapi-gin/controllers/productcontroller"
	"github.com/im4mbukh4ri/go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	models.ConnectDatabase()

	route.GET("/api/products", productcontroller.Index)
	route.GET("/api/product/:id", productcontroller.Show)
	route.POST("/api/product", productcontroller.Store)
	route.PUT("/api/product/:id", productcontroller.Update)
	route.DELETE("/api/product", productcontroller.Delete)

	route.Run()
}
