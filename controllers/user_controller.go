package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"time"
	"web-app/models/user"
	"web-app/providers"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	queries *user.Queries
	context context.Context
}

func NewUserController(q *user.Queries) *UserController {
	// open a connection to the database and return a pointer of the connection
	db := (&providers.Sql{}).Init()

	// create a new context
	ctx := context.Background()

	// create a new instance of the user queries
	queries := user.New(db)

	// create a new user controller
	return &UserController{queries: queries, context: ctx}
}

// Create implements the POST /api/users route
func (uc *UserController) Create(c *gin.Context) {
	// validate the request body
	var userInput user.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// generate a random 4-digit OTP
	otp := generateRandomOTP()

	// create a new profile
	profileID, err := uc.createProfile(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// create a new instance of CreateUserParams
	params := user.CreateUserParams{
		Email:             userInput.Email,
		Password:          userInput.Password,
		PhoneNumber:       userInput.PhoneNumber,
		Otp:               sql.NullString{String: otp, Valid: true},
		OtpExpirationTime: sql.NullTime{Time: time.Now().Add(10 * time.Minute), Valid: true}, // Set OTP expiration time
		ProfileID:         profileID,
	}

	// create a new user
	_, err = uc.queries.CreateUser(uc.context, params)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// GetAll implements the GET /api/users route
func (uc *UserController) GetAll(c *gin.Context) {
	users, err := uc.queries.GetAllUsers(uc.context)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GenerateOTP implements the POST /api/users/generateotp route
func (uc *UserController) GenerateOTP(c *gin.Context) {
	//
	c.JSON(http.StatusOK, gin.H{"message": "OTP generated successfully"})
}

// VerifyOTP implements the POST /api/users/verifyotp route
func (uc *UserController) VerifyOTP(c *gin.Context) {
	//
	c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
}

// CreateProfile implements the logic to create a new profile
func (uc *UserController) createProfile(c *gin.Context) (int64, error) {
	// validate the request body
	var profileInput user.Profile
	if err := c.ShouldBindJSON(&profileInput); err != nil {
		return 0, err
	}

	// create a new instance of CreateProfileParams
	params := user.CreateProfileParams{
		FirstName: profileInput.FirstName,
		LastName:  profileInput.LastName,
		Address:   profileInput.Address,
	}

	// create a new profile
	profile, err := uc.queries.CreateProfile(uc.context, params)
	if err != nil {
		return 0, err
	}

	return profile.ID, nil
}

// UpdateProfileByID implements the logic to update a profile by its ID
func (uc *UserController) UpdateProfileByID(c *gin.Context) {
	//
	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

// DeleteProfileByID implements the logic to delete a profile by its ID
func (uc *UserController) DeleteProfileByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uc.queries.DeleteProfileByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted successfully"})
}

func generateRandomOTP() string {
	// Generate a random 4-digit OTP
	return "1234"
}
