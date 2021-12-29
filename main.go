package main

import (
	"simpleapi_go/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/products", getProducts)

	router.Run("localhost:8083")
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.GetProducts())
}
