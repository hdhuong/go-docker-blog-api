package main

import (
	"blog/api/controllers"
	"blog/api/repositories"
	"blog/api/routes"
	"blog/api/services"
	"blog/config"
	"blog/database"
	"blog/models"
	"os"
)

func init() {
	config.LoadEnv()
}

func main() {

	router := config.NewGinRouter()
	db := database.ConnectDatabase()
	postRepository := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepository)
	postController := controllers.NewPostController(postService)
	postRoute := routes.NewPostRoute(postController, router)
	postRoute.Setup()

	db.DB.AutoMigrate(&models.Post{})
	
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000" // Default to port 8000 to match .env.example
	}
	router.Gin.Run(":" + port)
}
