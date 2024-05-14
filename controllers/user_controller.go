package controllers

import (
	"net/http"
	"web-app/models"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (*UserController) Index(ctx *gin.Context) {
	// get all users
	users, err := (&models.User{}).GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"users":   users,
	})
}

func (*UserController) Create(ctx *gin.Context) {
	var createUser models.CreateUser

	if err := ctx.ShouldBindJSON(&createUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create a user
	userID, err := (&models.User{}).Create(createUser)

	if err != nil {
		if userID == 422 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"user_id": userID,
	})
}

func (*UserController) GenerateOTP(ctx *gin.Context) {
	var user models.GenerateOTP

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// generate otp
	code, err := (&models.User{}).GenerateOTP(user.Phone)

	if err != nil {
		if code != 0 {
			ctx.JSON(code, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (*UserController) VerifyOTP(ctx *gin.Context) {
	var user models.VerifyOTP

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// verify otp
	code, err := (&models.User{}).VerifyOTP(user)

	if err != nil {
		if code != 0 {
			ctx.JSON(code, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
