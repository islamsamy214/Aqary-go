package providers

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// make a service provider for the database
type Sql struct{}

func (*Sql) Init() *sql.DB {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// connect to the pgsql database
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUsername, dbPassword, dbName)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic("Could not conncet to database due to: " + err.Error())
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	fmt.Println("session opened")
	return db
}
