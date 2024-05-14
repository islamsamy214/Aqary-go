package bootstrap

import (
	"web-app/migrations"
	"web-app/routes"

	"github.com/gin-gonic/gin"
)

func Boot() {
	// run the migrations
	(&migrations.ProfileTable{}).CreateTables()
	(&migrations.UserTable{}).CreateTables()

	server := gin.Default() // this craetes the server or engine

	routes.Web(server) // regester the web routes

	server.Run("0.0.0.0:8000") // start the server
}
