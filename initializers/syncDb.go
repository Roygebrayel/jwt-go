package initializers

import "github.com/Roygebrayel/jwt-go/models"

func SyncDb(){
	DB.AutoMigrate(&models.User{})
}