package initializers

import "github.com/ggomes061/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
