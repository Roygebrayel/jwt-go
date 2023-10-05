package initializers

import "jwt-go/models"

func SyncDb(){
	DB.AutoMigrate(&models.User{})
}