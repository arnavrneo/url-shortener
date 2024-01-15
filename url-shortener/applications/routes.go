package applications

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

	router.GET("/", handleForm)
	router.POST("/shorten", handleShorten)
	router.GET("/short/:id", handleRedirect)
	router.POST("/signup", SignUp)
	router.POST("/login", Login)
	router.GET("/validate", middleware.RequireAuth, Validate) // middleware for protecting routes

	router.Run()

	return router
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
