package migrations

import (
	"pedy/models"

	"gorm.io/gorm"
)

func RunAutoMigrations(db *gorm.DB) {

	db.AutoMigrate(models.User{}, models.Restaurant{})
}
