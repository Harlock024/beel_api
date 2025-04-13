package main

import (
	"beel_api/src/api/middleware"
	"beel_api/src/api/routes"
	"beel_api/src/db"
	"beel_api/src/migrations"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	migrations.Run()
	r := gin.Default()

	auth := r.Group("/auth")
	{
		routes.AuthRoutes(auth)
	}

	api := r.Group("/api")
	{

		api.Use(middleware.AuthMiddleware())

		routes.TaskRoutes(api)
		routes.ListRoutes(api)
		routes.TagRoutes(api)

	}

	r.Run("localhost:8080")
}
