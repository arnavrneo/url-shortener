package applications

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"url-shortener/initializers"
	"url-shortener/models"
)

func SignUp(c *gin.Context) {
	// Get the email/pass off req body
	var body struct {
		Name     string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "failed to read the body",
			})

		return
	}

	// Hash the pass
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	// Create the user
	coll := initializers.Client.Database("user_db").Collection("user_db_urlshortener")
	newSignUp := models.UserModel{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hash),
	}
	result, err := coll.InsertOne(c, newSignUp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot sign you up",
		})
		return
	}
	defer func() {
		if err = initializers.Client.Disconnect(c); err != nil {
			panic(err)
		}
	}()

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"success, objectID: ": result.InsertedID,
	})
}
