package routes

import (
	"jwt-go/controllers"
	"jwt-go/middleware"

	"github.com/gin-gonic/gin"
)


func SetupRoutes(router *gin.Engine) {
	
router.POST("/signup",  controllers.Signup)
  router.POST("/login", controllers.Login)
  router.GET("/validate",middleware.RequireAuth, controllers.Validate)
}

