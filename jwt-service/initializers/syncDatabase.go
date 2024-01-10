package initializers

import "jwt-service/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
