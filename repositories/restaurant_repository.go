package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"pedy/models"
)

type RestaurantRepository struct {
	DB *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) RestaurantRepository {
	return RestaurantRepository{DB: db}
}
func (r RestaurantRepository) All() ([]models.Restaurant) {
	var restaurants []models.Restaurant
	fmt.Println(r.DB)
	r.DB.Find(&restaurants)
	return restaurants
}
