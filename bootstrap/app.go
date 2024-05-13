package bootstrap

import (
	"web-app/routes"

	"github.com/gin-gonic/gin"
)

func Boot() {
	server := gin.Default() // this craetes the server or engine

	routes.Web(server) // regester the web routes

	server.Run("0.0.0.0:8000") // start the server
}
