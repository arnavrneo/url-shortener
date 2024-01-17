package applications

import (
	"github.com/gin-gonic/gin"
	"log"
	"url-shortener/middleware"
)

func loadRoutes() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", loginPage)            // html serving
	router.GET("/signupUser", signupPage) // html serving
	router.POST("/login", login)
	router.POST("/signup", signUp)
	router.GET("/main", middleware.RequireAuth, handleForm)        // html serving
	router.POST("/shorten", middleware.RequireAuth, handleShorten) // TODO: whether to merge this under /main
	router.GET("/short/:id", handleRedirect)
	router.GET("/validate", middleware.RequireAuth, validate) // middleware for protecting routes

	err := router.Run()
	if err != nil {
		log.Fatal("error running the router.")
	}
	return router
}
