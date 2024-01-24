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
	shortenedURL := fmt.Sprintf("http://localhost:%s/api/short/%s", os.Getenv("PORT"), shortKey)

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
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no short key found",
		})
		return
	}

	if originalURL == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no short key found",
		})
		return
	}

	if shortKey == urls.ShortKey {
		originalURL = urls.OriginalURL
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}
