package controllers

import (
	"fmt"
	"jwt-go/initializers"
	"jwt-go/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup (c *gin.Context) {
	// req body

	var body struct {
		Email string
		Password string 
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error" : "no body parsed",
		})
		return
	}



	// hash pass

	hash , err := bcrypt.GenerateFromPassword([]byte(body.Password),10)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error" : "no body parsed",
		})
		return

	}
	// save 

	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
	}

	// respond

	c.JSON(http.StatusAccepted,gin.H{
		"entries" : "everything okay",
	})


}

func Login( c *gin.Context){
	// get the email and pass from body req

	var body struct {
		Email string
		Password string 
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error" : "no body parsed",
		})
		return
	}

	//lookup request user

	var users models.User;
	initializers.DB.First(&users, "email = ?", body.Email)

	if users.ID == 0 {
			
		c.JSON(http.StatusBadRequest,gin.H{
			"error" : "invalid email or pass",
		})
		return
	}
	
	


	// compare the pass with the hashed passwrod
err := bcrypt.CompareHashAndPassword([]byte(users.Password),[]byte(body.Password))

if err != nil {
c.JSON(http.StatusBadRequest,gin.H{
	"eror" : "there is an error",
})
return
}
	
	// direct the user

	// generate jwt token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"sub" : users.ID,
		"exp" : time.Now().Add(time.Hour* 24 * 30).Unix(),


	})
	tokenString, err := token.SignedString([]byte (os.Getenv("SECRET")))
	if err!=nil {
		fmt.Println(err);

	}

	// send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("authorization" ,tokenString,3600 * 24 * 30 ,"" ,"",false,true )
	c.JSON(http.StatusOK,gin.H{})
	
}