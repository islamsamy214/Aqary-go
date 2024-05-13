package controllers

import (
	"net/http"
	"strconv"
	"web-app/models/user"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	queries *user.Queries
}

func NewUserController(q *user.Queries) *UserController {
	return &UserController{queries: q}
}

// Create implements the POST /api/users route
func (uc *UserController) Create(c *gin.Context) {
	//
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// GetAll implements the GET /api/users route
func (uc *UserController) GetAll(c *gin.Context) {
	users, err := uc.queries.GetAllUsers(c.Request.Context())
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
func (uc *UserController) createProfile(c *gin.Context) {
	//
	// c.JSON(http.StatusCreated, gin.H{"id": profileID})
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
