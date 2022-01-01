package main

import (
	"simpleapi_go/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/products", getProducts)

	router.GET("/product/:code", getProduct)

	router.POST("/products", addProduct)

	router.Run("localhost:8083")
}

func getProducts(c *gin.Context) {

	products := models.GetProducts()

	if products == nil || len(products) == 0 {

		c.AbortWithStatus(http.StatusNotFound)

	} else {

		c.IndentedJSON(http.StatusOK, products)

	}
}

func getProduct(c *gin.Context) {

	code := c.Param("code")

	product := models.GetProduct(code)

	if product == nil {
		c.AbortWithStatus(http.StatusNotFound)

	} else {

		c.IndentedJSON(http.StatusOK, product)

	}

}

func addProduct(c *gin.Context) {

}
