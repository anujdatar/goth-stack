package main

import (
	"log"
	"net/http"
	"os"

	dbController "github.com/anujdatar/goth-stack/controllers/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error. Unable to load .env file: ", err)
	}
	PORT := os.Getenv("PORT")

	db := dbController.ConnectToDb()
	defer db.Close()
	dbController.CreateDbTables(db)

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
