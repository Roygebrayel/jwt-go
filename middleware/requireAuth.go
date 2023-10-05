package middleware

import (
	"fmt"
	"github.com/Roygebrayel/jwt-go/initializers"
	"github.com/Roygebrayel/jwt-go/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth (c *gin.Context){
	// get the cookie from req

	tokenString , err := c.Cookie("authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	

token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// Don't forget to validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	return []byte (os.Getenv("SECRET")), nil
})

if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	if float64 (time.Now().Unix()) > claims["exp"].(float64) {
		c.JSON(201,"expired")
	}

	var users models.User
	initializers.DB.First(&users,claims["sub"])

	if users.ID == 0 {
		c.JSON(503 ,"user not found ")
	}
	c.Set("user",users)
} else {
	c.AbortWithStatus(http.StatusUnauthorized)
}

	
	c.Next();
}