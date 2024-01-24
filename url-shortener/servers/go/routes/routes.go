package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"url-shortener/controllers"
	"url-shortener/middleware"
)

func LoadRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) // for production release
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3000/signup", "http://localhost:3000/main"}, // in production, allow only fronted production endpoint
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))

	router.POST("/api/login", applications.Login)
	router.POST("/api/logout", applications.Logout)
	router.POST("/api/register", applications.Register)
	router.POST("/api/shorten", middleware.RequireAuth, applications.HandleShorten) // TODO: whether to merge this under /main
	router.GET("/api/short/:id", applications.HandleRedirect)
	router.GET("/api/user", middleware.RequireAuth, applications.GetUser)

	return router
}
