package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/libsql/libsql-client-go/libsql"
)

var Database *sql.DB

func Connect() {
	var DB_URL string
	var DB_String string
	if os.Getenv("APP_ENV") == "prod" {
		DB_URL = os.Getenv("DB_URL")
		DB_AUTH_TOKEN := os.Getenv("TURSO_AUTH_TOKEN")
		if DB_AUTH_TOKEN == "" {
			log.Fatal("Error. DB_AUTH_TOKEN not set")
		}
		DB_String = DB_URL + "?authToken=" + DB_AUTH_TOKEN
	} else {
		DB_URL = os.Getenv("DB_URL")
		DB_String = DB_URL
	}

	db, err := sql.Open("libsql", DB_String)
	if err != nil {
		log.Fatal("Error. Unable to connect to database: ", err)
	}
	log.Println("Connected to database at: ", DB_URL)
	Database = db
}
