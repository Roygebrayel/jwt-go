package main

import (
	"github.com/Roygebrayel/jwt-go/initializers"
	"github.com/Roygebrayel/jwt-go/routes"

	"github.com/gin-gonic/gin"
)

func init (){
	initializers.LoadEnvVariables();
	initializers.ConnectToDB();
  initializers.SyncDb();
}
func main (){
	
  r := gin.Default()

  routes.SetupRoutes(r);


  r.Run() 
}