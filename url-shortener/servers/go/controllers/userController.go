package controllers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
	"url-shortener/models"
)

type reqBody struct {
	Username string `form:"username" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// Register the user
func Register(c *gin.Context) {
	// Get the email/pass off req body
	var body reqBody

	if c.ShouldBind(&body) != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "failed to read the request",
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
	Client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot connect to the client",
		})
	}

	coll := Client.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("DATABASE_COLLECTION"))
	newSignUp := models.UserModel{
		Username: body.Username,
		Email:    body.Email,
		Password: string(hash),
	}

	userExists := checkUser(coll, body) // checks if the email already exists

	if userExists == false {
		_, err = coll.InsertOne(c, newSignUp)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "cannot sign you up; bad request",
			})
			return
		}
		defer func() {
			if err = Client.Disconnect(c); err != nil {
				panic(err)
			}
		}()
		c.JSON(http.StatusOK, gin.H{
			"msg": "register request successful",
		})
	}
}

// Login logs in the user
func Login(c *gin.Context) {
	// Get the email and pass off the req body
	var body reqBody

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "failed to read the request",
			})
		return
	}

	// look up the req user
	Client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot connect to the client",
		})
	}

	coll := Client.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("DATABASE_COLLECTION"))
	filter := bson.D{{"email", body.Email}}

	var result models.UserModel
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusUnauthorized,
				gin.H{
					"error": "invalid credentials",
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

	// SEND IT BACK
	// as a cookie
	//c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"Authorization",
		tokenString,
		3600,
		"",
		"",
		false,
		true)

	c.JSON(http.StatusOK, gin.H{
		"message": "login request successful",
	})
	// as a jwt token to be stored in the storage session or whatever
	//c.JSON(http.StatusOK, gin.H{
	//	"token": tokenString,
	//})
}

// checkUser checks whether the user already exists or not
func checkUser(coll *mongo.Collection, body reqBody) bool {
	var result models.UserModel
	emailFilter := bson.D{{"email", body.Email}}

	emailErr := coll.FindOne(context.TODO(), emailFilter).Decode(&result)

	if emailErr != nil {
		return false
	} else {
		return true
	}
}

// Logout logs out the user
func Logout(c *gin.Context) {
	c.SetCookie(
		"Authorization",
		"",
		0,
		"",
		"",
		false,
		true)

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully logged out",
	})
}

// GetUser fetch user detail
func GetUser(c *gin.Context) {
	user, err := c.Get("user")
	if err == false {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"username": user,
	})
}
