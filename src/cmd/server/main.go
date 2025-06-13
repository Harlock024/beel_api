package main

import (
	"beel_api/src/api/handlers"
	"beel_api/src/api/middleware"
	"beel_api/src/api/routes"
	"beel_api/src/db"
	"beel_api/src/internal/repositories"
	"beel_api/src/internal/services"
	"beel_api/src/migrations"

	//	"beel_api/src/migrations"

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
		authHandler := handlers.NewAuthHandler(services.NewAuthService(repositories.NewUserRepository(db.DB), repositories.NewRefreshRepository(db.DB)))
		routes.AuthRoutes(auth, authHandler)
	}
	api := r.Group("/api")
	{
		// Initialize repositories and services
		taskHandler := handlers.NewTaskHandler(services.NewTaskService(repositories.NewTaskRepository(db.DB)))
		listHandler := handlers.NewListHandler(services.NewListService(*repositories.NewListRepository(db.DB)))
		tagHandler := handlers.NewTagHandler(services.NewTagService(repositories.NewTagRepository(db.DB)))

		api.Use(middleware.AuthMiddleware())
		routes.RefreshRoutes(api)
		routes.TaskRoutes(api, taskHandler)

		routes.ListRoutes(api, listHandler)
		routes.TagRoutes(api, tagHandler)
	}

	r.Run("localhost:8080")
}
