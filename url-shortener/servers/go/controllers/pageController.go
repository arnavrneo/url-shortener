package applications

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"url-shortener/utils"
)

type urlMap struct {
	ShortenedURL string
	OriginalURL  string
	ShortKey     string
}

var urls urlMap

func HandleForm(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "form.html", nil)
}

func HandleShorten(c *gin.Context) {
	originalURL := c.PostForm("url")
	if originalURL == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	shortKey := utils.GenerateShortKey() // TODO: check for duplicate keys
	shortenedURL := fmt.Sprintf("http://localhost:%s/short/%s", os.Getenv("PORT"), shortKey)

	urls = urlMap{
		ShortenedURL: shortenedURL,
		OriginalURL:  originalURL,
		ShortKey:     shortKey,
	}

	c.JSON(http.StatusOK, gin.H{
		"shorten_link": urls.ShortenedURL,
	})
}

func HandleRedirect(c *gin.Context) {
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
