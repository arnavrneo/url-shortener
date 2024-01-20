package applications

import (
	"github.com/gin-gonic/gin"
	"url-shortener/middleware"
)

func LoadRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) // for production release
	router := gin.Default()

	router.LoadHTMLGlob("templates/*") // ../templates/* for testing

	router.GET("/", loginPage)            // html serving
	router.GET("/signupUser", signupPage) // html serving
	router.POST("/login", Login)
	router.POST("/signup", signUp)
	router.GET("/main", middleware.RequireAuth, handleForm)        // html serving
	router.POST("/shorten", middleware.RequireAuth, handleShorten) // TODO: whether to merge this under /main
	router.GET("/short/:id", handleRedirect)
	router.GET("/validate", middleware.RequireAuth, validate) // middleware for protecting routes

	return router
}