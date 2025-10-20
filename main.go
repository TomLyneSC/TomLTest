package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Healthcheck endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	return r
}

func main() {
	r := setupRouter()

	fmt.Println("hello look at me!")

	r.Run(":8080")
}
