package providers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// make a service provider for the database
type Sql struct{}

func (*Sql) Init() *sql.DB {
	// connect to the pgsql database
	db, err := sql.Open("postgres", "postgres://olo:password@localhost:5432/aqary?sslmode=disable")

	if err != nil {
		panic("Could not conncet to database due to: " + err.Error())
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("session opened")
	return db
}
