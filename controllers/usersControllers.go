package controllers

import (
	"jwt-go/initializers"
	"jwt-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
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