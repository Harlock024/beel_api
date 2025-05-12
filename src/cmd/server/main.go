package main

import (
	"beel_api/src/api/middleware"
	"beel_api/src/api/routes"
	"beel_api/src/db"
	"beel_api/src/migrations"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	migrations.Run()
	r := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:4321"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(config))

	auth := r.Group("/auth")
	{
		routes.AuthRoutes(auth)
	}

	api := r.Group("/api")
	{
		api.Use(middleware.AuthMiddleware())
		routes.RefreshRoutes(api)
		routes.TaskRoutes(api)
		routes.ListRoutes(api)
		routes.TagRoutes(api)

	}

	r.Run("localhost:8080")
}
