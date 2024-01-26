package applications

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"url-shortener/initializers"
	"url-shortener/utils"
)

type urlMap struct {
	ShortenedURL string
	OriginalURL  string
	ShortKey     string
}

var urls urlMap

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

	rdb := initializers.ConnectRedis()

	// set in the redis
	err := rdb.Set(initializers.Ctx, urls.ShortKey, urls.OriginalURL, 0).Err()
	if err != nil {
		fmt.Println("cannot set the values in redis")
	}

	c.JSON(http.StatusOK, gin.H{
		"shorten_link": urls.ShortenedURL,
	})
}

func HandleRedirect(c *gin.Context) {
	var originalURL string
	shortKey := c.Param("id")

	if shortKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}

	rdb := initializers.ConnectRedis()
	// get the corresponding original url
	originalURL, err := rdb.Get(initializers.Ctx, urls.ShortKey).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "shortKey not found",
		})
	}

	if originalURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}
