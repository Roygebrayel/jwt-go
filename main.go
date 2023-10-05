package main

import (
	
	"jwt-go/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init (){
	initializers.LoadEnvVariables();
	initializers.ConnectToDB();
  initializers.SyncDb();
}
func main (){
	
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() 
}