package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"url-shortener/controllers"
	"url-shortener/middleware"
)

func LoadRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) // for production release
	router := gin.Default()

	allowedOrigins := os.Getenv("ORIGINS")

	router.Use(cors.New(cors.Config{
		//AllowOrigins: allowedOrigins,
		AllowOrigins:     []string{allowedOrigins}, // in production, allow only fronted production endpoint
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))

	router.POST("/api/login", controllers.Login)
	router.POST("/api/logout", controllers.Logout)
	router.POST("/api/register", controllers.Register)
	router.POST("/api/shorten", middleware.RequireAuth, controllers.HandleShorten) // TODO: whether to merge this under /main
	router.GET("/short/:id", controllers.HandleRedirect)
	router.GET("/api/user", middleware.RequireAuth, controllers.GetUser)

	return router
}
