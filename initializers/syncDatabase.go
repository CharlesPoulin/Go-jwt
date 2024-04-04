package initializers

import "Go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})

}
