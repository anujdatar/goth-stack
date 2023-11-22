package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	fmt.Println("Server is running on port 8080")
	fmt.Println("Connect using http://localhost:8080 or http://0.0.0.0:8080")

	r.Run()
}
