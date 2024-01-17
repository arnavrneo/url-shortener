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

type reqBody struct {
	Username string `form:"username" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// SignUp the user
func signUp(c *gin.Context) {
	// Get the email/pass off req body
	var body reqBody

	if c.ShouldBind(&body) != nil {
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
	coll := initializers.Client.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("DATABASE_COLLECTION"))
	newSignUp := models.UserModel{
		Username: body.Username,
		Email:    body.Email,
		Password: string(hash),
	}

	userExists, value := checkUser(coll, body) // checks if the username already exists

	if userExists == false {
		_, err = coll.InsertOne(c, newSignUp)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "cannot sign you up; bad request",
			})
			return
		}
		defer func() {
			if err = initializers.Client.Disconnect(c); err != nil {
				panic(err)
			}
		}()
		// Respond & Redirect
		//c.JSON(http.StatusOK, gin.H{
		//	"success, objectID: ": result.InsertedID,
		//})
		c.Redirect(http.StatusFound, "/")
	} else if value == "email" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "email already exists",
		})
	} else if value == "username" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "username already exists",
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
	coll := initializers.Client.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("DATABASE_COLLECTION"))
	filter := bson.D{{"email", body.Email}}

	var result models.UserModel
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusBadRequest,
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
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600, "", "", false, true)

	// as a jwt token to be stored in the storage session or whatever
	//c.JSON(http.StatusOK, gin.H{
	//	"token": tokenString,
	//})

	c.Redirect(http.StatusFound, "/main") // cannot redirect POST to GET route
	c.Abort()                             // Aborts the pending handlers
}

// Validate helper function
func validate(c *gin.Context) {
	user, err := c.Get("user")
	if err != false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

// checkUser checks whether the user already exists or not
func checkUser(coll *mongo.Collection, body reqBody) (bool, string) {
	var result models.UserModel
	usernameFilter := bson.D{{"username", body.Username}}
	emailFilter := bson.D{{"email", body.Email}}

	usernameErr := coll.FindOne(context.TODO(), usernameFilter).Decode(&result) // error will be returned if user is not present
	emailErr := coll.FindOne(context.TODO(), emailFilter).Decode(&result)

	// username no error exists; email no error duplicate
	if usernameErr != nil {
		if emailErr != nil {
			return false, "" // user doesnt exists
		} else {
			return true, "email" // username unique; email exists
		}
	} else {
		return true, "username" // username exists
	}
}
