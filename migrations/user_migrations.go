package migrations

import (
	"fmt"
	"web-app/providers"
)

type UserTable struct{}

func (*UserTable) CreateTables() {
	db := (&providers.Sql{}).Init()

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id BIGSERIAL PRIMARY KEY,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			phone_number VARCHAR(20) UNIQUE NOT NULL,
			otp VARCHAR(6),
			otp_expiration_time TIMESTAMP,
			profile_id INT,
			FOREIGN KEY (profile_id) REFERENCES profiles(id)
		);
	`)

	if err != nil {
		panic("Could not create users table due to: " + err.Error())
	}

	fmt.Println("users table created")
	defer db.Close()
}
