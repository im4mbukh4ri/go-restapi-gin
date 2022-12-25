package productcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/im4mbukh4ri/go-restapi-gin/models"
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
