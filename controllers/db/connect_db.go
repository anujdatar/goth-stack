package controllers

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/libsql/libsql-client-go/libsql"
)

func ConnectToDb() *sql.DB {
	var DB_URL string
	var DB_String string
	if os.Getenv("APP_ENV") == "prod" {
		DB_URL = os.Getenv("TURSO_URL")
		DB_AUTH_TOKEN := os.Getenv("TURSO_AUTH_TOKEN")
		if DB_AUTH_TOKEN == "" {
			log.Fatal("Error. TURSO_AUTH_TOKEN not set")
		}
		DB_String = DB_URL + "?authToken=" + DB_AUTH_TOKEN
	} else {
		DB_URL = os.Getenv("DB_URL")
		DB_String = DB_URL
	}
	// DB_URL := "http://localhost:8080"
	log.Println("DB_URL: ", DB_String)

	db, err := sql.Open("libsql", DB_URL)
	if err != nil {
		log.Fatal("Error. Unable to connect to database: ", err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("Error. Unable to ping database: ", pingErr)
	}
	log.Println("Connected to database at: ", DB_URL)

	return db
}
