package main

import (
	"beel_api/src/api/handlers"
	"beel_api/src/api/middleware"
	"beel_api/src/api/routes"
	"beel_api/src/db"
	"beel_api/src/internal/repositories"
	"beel_api/src/internal/services"
	"os"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	config := cors.Config{
		AllowOrigins:  []string{"http://localhost:4321"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}
	r.Use(cors.New(config))

	auth := r.Group("/auth")
	{
		authHandler := handlers.NewAuthHandler(services.NewAuthService(repositories.NewUserRepository(db.DB), repositories.NewRefreshRepository(db.DB)))
		routes.AuthRoutes(auth, authHandler)
		refreshHandler := handlers.NewRefreshHandler(services.NewRefreshServices(repositories.NewRefreshRepository(db.DB), repositories.NewUserRepository(db.DB)))
		routes.RefreshRoutes(auth, refreshHandler)
	}
	api := r.Group("/api")
	{
		taskHandler := handlers.NewTaskHandler(services.NewTaskService(repositories.NewTaskRepository(db.DB)))
		listHandler := handlers.NewListHandler(services.NewListService(repositories.NewListRepository(db.DB)))
		tagHandler := handlers.NewTagHandler(services.NewTagService(repositories.NewTagRepository(db.DB)))
		api.Use(middleware.AuthMiddleware())

		routes.TaskRoutes(api, taskHandler)
		routes.ListRoutes(api, listHandler)
		routes.TagRoutes(api, tagHandler)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
