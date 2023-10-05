package main

import (
	"jwt-go/controllers"
	"jwt-go/initializers"
	

	"github.com/gin-gonic/gin"
)

func init (){
	initializers.LoadEnvVariables();
	initializers.ConnectToDB();
  initializers.SyncDb();
}
func main (){
	
  r := gin.Default()
  r.POST("/signup",  controllers.Signup)
  r.Run() 
}