package main

import (
	"simpleapi_go/models"

	"net/http"

	"github.com/gin-gonic/gin"

	"strconv"
)

func main() {

	router := gin.Default()

	router.GET("/products/:offset/:limit", getProducts)

	router.Run("localhost:8083")
}

func getProducts(c *gin.Context) {

	offset := c.Param("offset")

	limit := c.Param("limit")

	offset_int, _ := strconv.Atoi(offset)
	limit_int, _ := strconv.Atoi(limit)

	c.IndentedJSON(http.StatusOK, models.GetProducts(offset_int, limit_int))
}
