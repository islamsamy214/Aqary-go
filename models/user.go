package models

import (
	"errors"
	"time"
	"web-app/providers"
)

type User struct {
	ID        int    `json:"id"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Otp       int    `json:"otp"`
	OtpExpiry string `json:"otp_expiry"`
	ProfileID int    `json:"profile_id"`
}

type CreateUser struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Address   string `json:"address" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	Phone     string `json:"phone" binding:"required"`
}

type GenerateOTP struct {
	Phone string `json:"phone" binding:"required"`
}

type VerifyOTP struct {
	Phone string `json:"phone" binding:"required"`
	Otp   int    `json:"otp" binding:"required"`
}

func (u *User) GetAll() ([]User, error) {
	var users []User

	db := (&providers.Sql{}).Init()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return users, err
	}
	defer rows.Close() // Make sure to close the rows when done

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Password, &user.Email, &user.Phone, &user.Otp, &user.OtpExpiry, &user.ProfileID)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	// Check for any errors that may have occurred during iteration
	if err = rows.Err(); err != nil {
		return users, err
	}

	// get the profile for each user
	for i, user := range users {
		profile, err := (&Profile{ID: user.ProfileID}).GetProfile()
		if err != nil {
			return users, err
		}
		users[i].Password = profile.FirstName
		users[i].Email = profile.LastName
		users[i].Phone = profile.Address
	}

	return users, nil
}

func (u *User) Create(user CreateUser) (int, error) {
	db := (&providers.Sql{}).Init()

	var userID int
	var profileID int

	// Check if the phone number already exists in the database
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE phone = $1", user.Phone).Scan(&count)
	if err != nil || count > 0 {
		// new error message
		return 422, errors.New("phone number already exists")
	}

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback() // Rollback the transaction if there's an error

	// Create a profile and get the profile ID
	err = tx.QueryRow("INSERT INTO profiles (first_name, last_name, address) VALUES ($1, $2, $3) RETURNING id", user.FirstName, user.LastName, user.Address).Scan(&profileID)
	if err != nil {
		return 0, err
	}

	// Create a user and get the user ID
	err = tx.QueryRow("INSERT INTO users (password, email, phone, otp, otp_expiry, profile_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", user.Password, user.Email, user.Phone, 0, time.Now(), profileID).Scan(&userID)
	if err != nil {
		return 0, err
	}

	// Commit the transaction if everything is successful
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (u *User) GenerateOTP(phone string) (int, error) {
	db := (&providers.Sql{}).Init()

	var otp int

	// Check if the phone number exists in the database
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE phone = $1", phone).Scan(&count)
	if err != nil || count == 0 {
		// new error message
		return 404, errors.New("phone number does not exist")
	}

	// generate a random 6 digit number
	otp = 123456

	_, err = db.Exec("UPDATE users SET otp = $1, otp_expiry = $2 WHERE phone = $3", otp, time.Now().Add(time.Minute), phone)
	if err != nil {
		return 0, err
	}

	return otp, nil
}

func (u *User) VerifyOTP(verify VerifyOTP) (int, error) {
	db := (&providers.Sql{}).Init()

	// Check if the phone number exists in the database
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE phone = $1", verify.Phone).Scan(&count)
	if err != nil || count == 0 {
		// new error message
		return 404, errors.New("phone number does not exist")
	}

	// Check if the OTP is correct
	var otp int
	var otpExpiry time.Time
	err = db.QueryRow("SELECT otp, otp_expiry FROM users WHERE phone = $1", verify.Phone).Scan(&otp, &otpExpiry)
	if err != nil {
		return 0, err
	}

	if otp != verify.Otp {
		return 422, errors.New("incorrect OTP")
	}

	if time.Now().After(otpExpiry) {
		return 422, errors.New("OTP expired")
	}

	_, err = db.Exec("UPDATE users SET otp = $1, otp_expiry = $2 WHERE phone = $3", 0, time.Now(), verify.Phone)
	if err != nil {
		return 0, err
	}

	return 0, nil
}
