package main

import (
	"blog/api/controllers"
	"blog/api/repositories"
	"blog/api/routes"
	"blog/api/services"
	"blog/config"
	"blog/database"
	"blog/models"
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
	router.Gin.Run(":8001")
}
