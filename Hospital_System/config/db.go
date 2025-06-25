package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("postgres", "host=localhost port=6666 user=postgres password=database@123 dbname=hospital_db sslmode=disable")
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL!")
}
