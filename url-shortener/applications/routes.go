package applications

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"url-shortener/middleware"
	"url-shortener/utils"
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
	router.POST("/login", Login)
	router.POST("/signup", SignUp)
	router.GET("/main", middleware.RequireAuth, handleForm)
	router.POST("/shorten", handleShorten) // TODO: protect this and the below path; merge under /main
	router.GET("/short/:id", handleRedirect)
	router.GET("/validate", middleware.RequireAuth, Validate) // middleware for protecting routes

	err := router.Run()
	if err != nil {
		log.Fatal("error running the router.")
	}
	return router
}

func loginPage(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "login.html", nil)
}

func signupPage(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "signup.html", nil)
}

func handleForm(c *gin.Context) {
	// HTML Form
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "form.html", nil)
}

func handleShorten(c *gin.Context) {
	originalURL := c.PostForm("url")
	if originalURL == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	shortKey := utils.GenerateShortKey()
	shortenedURL := fmt.Sprintf("http://localhost:8000/short/%s", shortKey)

	urls = urlMap{
		ShortenedURL: shortenedURL,
		OriginalURL:  originalURL,
		ShortKey:     shortKey,
	}

	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "urlShorten.html", urls)
}

func handleRedirect(c *gin.Context) {
	var originalURL string
	shortKey := c.Param("id")

	if shortKey == "" {
		c.HTML(http.StatusUnauthorized, "error.html", nil)
		return
	}

	if shortKey == urls.ShortKey {
		originalURL = urls.OriginalURL
	}

	if originalURL == "" {
		c.HTML(http.StatusUnauthorized, "error.html", nil)
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}
