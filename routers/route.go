package routers

import (
	"example.id/mygram/controllers"
	"example.id/mygram/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	commentsRoute.GET("/", controllers.GetAllComments)
	commentsRoute.PUT("/:commentId", controllers.UpdateComment)
	commentsRoute.DELETE("/:commentId", controllers.DeleteComment)
	socmedsRoute := router.Group("socialmedias", middlewares.JwtAuthMiddleware())
	socmedsRoute.POST("/", controllers.CreateSocialMedia)
	socmedsRoute.GET("/", controllers.GetAllSocialMedias)
	socmedsRoute.PUT("/:socialMediaId", controllers.UpdateSocialMedia)
	socmedsRoute.DELETE("/:socialMediaId", controllers.DeleteSocialMedia)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
