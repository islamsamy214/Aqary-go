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

	db, err := sql.Open("postgres", "postgres://olo:password@localhost:5432/aqary-test?sslmode=disable")

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

// db := (&providers.Sql{}).Init() // initialize the database
// 	// excute sample query
// 	// 	INSERT INTO users (email, password, phone_number, profile_id)
// 	// VALUES ($1, $2, $3, $4)
// 	// RETURNING id;
// 	_, err := db.Exec("INSERT INTO profile (first_name, last_name, address) VALUES ($1, $2, $3)", "John", "Doe", "123 Main St")

// 	if err != nil {
// 		panic("Could not insert into database due to: " + err.Error())
// 	}
