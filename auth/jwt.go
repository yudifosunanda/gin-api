// auth/jwt.go
package auth

import (
	"time"
	"github.com/gin-gonic/gin"
	jwt "github.com/appleboy/gin-jwt/v2"
	"gin-api/models"
	"gin-api/db"
	"golang.org/x/crypto/bcrypt"
)

var identityKey = "user"

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// The middleware instance (exported so it can be reused)
var AuthMiddleware *jwt.GinJWTMiddleware

func InitJWT() {
	var err error
	AuthMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "api zone",
		Key:         []byte("secret jwt key here"), // put in env later
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals Login
			if err := c.ShouldBindJSON(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			var user models.User
			if err := db.DB.Where("username = ?", loginVals.Username).First(&user).Error; err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			
			if !CheckPasswordHash(loginVals.Password, user.Password){
				return nil, jwt.ErrFailedAuthentication
			}

			return &Login{Username: user.Username}, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// Always true for now (allow all logged-in users)
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"error": message})
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*Login); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &Login{
				Username: claims[identityKey].(string),
			}
		},
	})

	if err != nil {
		panic("JWT Error:" + err.Error())
	}
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}