package routes

import (
	"net/http"
	"web-app/controllers"

	"github.com/gin-gonic/gin"
)

func Web(route *gin.Engine) {
	route.GET("/", home)

	// user
	userController := controllers.UserController{}
	route.GET("/api/users", userController.Index)
	route.POST("/api/users", userController.Create)
	route.POST("/api/users/generateotp", userController.GenerateOTP)
	route.POST("/api/users/verifyotp", userController.VerifyOTP)

}

func home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "hello"})
}
