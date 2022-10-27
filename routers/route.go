package routers

import (
	"example.id/mygram/controllers"
	"example.id/mygram/middlewares"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.POST("users/register", controllers.RegisterUser)
	router.POST("users/login", controllers.LoginUser)
	router.PUT("users", middlewares.JwtAuthMiddleware(), controllers.UpdateUser)
	router.DELETE("users", middlewares.JwtAuthMiddleware(), controllers.DeleteUser)
	photosRoute := router.Group("photos", middlewares.JwtAuthMiddleware())
	photosRoute.POST("/", controllers.CreatePhoto)
	photosRoute.GET("/", controllers.GetAllPhotos)
	photosRoute.PUT("/:photoId", controllers.UpdatePhoto)
	photosRoute.DELETE("/:photoId", controllers.DeletePhoto)
	commentsRoute := router.Group("comments", middlewares.JwtAuthMiddleware())
	commentsRoute.POST("/", controllers.CreateComment)
	return router
}
