package applications

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
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

func Login(c *gin.Context) {
	// Get the email and pass off the req body
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

	// look up the req user
	coll := initializers.Client.Database("user_db").Collection("user_db_urlshortener")
	filter := bson.D{{"email", body.Email}}

	var result models.UserModel
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusBadRequest,
				gin.H{
					"error": "invalid email / password",
				})
			return
		}
		return
	}

	// compare sent in pass with saved user hash pass
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "wrong combination",
		})
		return
	}

	// generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": result.Email, // replace this with id (pref),
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	// sign and get the complete encoded token using the secret key (in .env)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})
	}

	// send it back
	// as a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600, "", "", false, true)

	// as a jwt token to be stored in the storage session or whatever
	//c.JSON(http.StatusOK, gin.H{
	//	"token": tokenString,
	//})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
