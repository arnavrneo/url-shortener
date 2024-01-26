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
	shortenedURL := fmt.Sprintf("http://localhost:%s/api/short/%s", os.Getenv("PORT"), shortKey)

	urls = urlMap{
		ShortenedURL: shortenedURL,
		OriginalURL:  originalURL,
		ShortKey:     shortKey,
	}

	rdb := initializers.ConnectRedis()

	// set in the redis
	err := rdb.Set(initializers.Ctx, urls.ShortKey, urls.OriginalURL, 0).Err()
	if err != nil {
		fmt.Printf("error %s", err)
	}

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"shorten_link": urls.ShortenedURL,
	})
}

func HandleRedirect(c *gin.Context) {
	//var originalURL string
	shortKey := c.Param("id")

	if shortKey == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	rdb := initializers.ConnectRedis()
	originalURL, err := rdb.Get(initializers.Ctx, urls.ShortKey).Result()
	fmt.Println("redirect reached ", originalURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "shortKey not found",
		})
	}

	if originalURL == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}
