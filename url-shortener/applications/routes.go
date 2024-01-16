package applications

import (
	"github.com/gin-gonic/gin"
	"log"
	"url-shortener/middleware"
)

type urlMap struct {
	ShortenedURL string
	OriginalURL  string
	ShortKey     string
}

var urls urlMap

func loadRoutes() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", loginPage)
	router.GET("/signupUser", signupPage)
	router.POST("/login", login)
	router.POST("/signup", signUp)
	router.GET("/main", middleware.RequireAuth, handleForm)
	router.POST("/shorten", handleShorten) // TODO: protect this and the below path; merge under /main
	router.GET("/short/:id", handleRedirect)
	router.GET("/validate", middleware.RequireAuth, validate) // middleware for protecting routes

	err := router.Run()
	if err != nil {
		log.Fatal("error running the router.")
	}
	return router
}
