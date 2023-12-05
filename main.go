package main

import (
	"log"
	"net/http"
	"os"

	"github.com/anujdatar/goth-stack/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatal("Error. Unable to load .env file: ", err)
	}
	PORT := os.Getenv("PORT")

	database.Connect()
	defer database.Database.Close()
	database.CreateDbTables(database.Database)

	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	log.Printf("Server is running on port %v\n", PORT)
	log.Printf("Connect using http://localhost:%v or http://0.0.0.0:%v\n", PORT, PORT)

	err = r.Run(":" + PORT)
	if err != nil {
		log.Fatal("Error. Unable to start server: ", err)
	}
}
