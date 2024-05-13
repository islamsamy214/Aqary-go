package migrations

import (
	"fmt"
	"web-app/providers"
)

type ProfileTable struct{}

func (*ProfileTable) CreateTables() {
	db := (&providers.Sql{}).Init()

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS profiles (
			id BIGSERIAL PRIMARY KEY,
			first_name VARCHAR(100),
			last_name VARCHAR(100),
			address TEXT
		);
	`)

	if err != nil {
		panic("Could not create profiles table due to: " + err.Error())
	}

	fmt.Println("profiles table created")
	defer db.Close()
}
