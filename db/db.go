package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func InitializeDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	dbdriver := os.Getenv("DBDRIVER")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")
	PORT := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=require", username, password, host, PORT, database)

	db, err = sql.Open(dbdriver, DBURL)

	if err != nil {
		log.Fatal("Error connecting to database:", err.Error())
	}

	createTable := `
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			price NUMERIC NOT NULL,
			stock smallint NOT NULL,
			created_at timestamptz DEFAULT now()
		)
	`

	_, err = db.Exec(createTable)

	if err != nil {
		log.Fatal("Error creating products table:", err.Error())
	}
	fmt.Println("Successfully connected to database")
}

func GetDB() *sql.DB {
	return db
}

func SetDB(dbVal *sql.DB) {
	db = dbVal
}
