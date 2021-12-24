package repositories

import (
	"errors"
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
	r.DB.Find(&restaurants)
	return restaurants
}

func (r RestaurantRepository) Find(id int) (models.Restaurant, error) {
	var restaurant models.Restaurant
	result := r.DB.Find(&restaurant, id)
	if result.Error != nil {
		return restaurant, result.Error
	}
	if result.RowsAffected == 0 {
		return restaurant, errors.New("No record found for provided id.")
	}
	return restaurant, nil
}

func (r RestaurantRepository) Create(data models.Restaurant) (models.Restaurant, error) {
	result := r.DB.Create(&data)
	if result.Error != nil {
		return data, result.Error
	}
	if result.RowsAffected == 0 {
		return data, errors.New("Could not insert data. 0 rows inserted.")
	}
	return data, nil
}

func (r RestaurantRepository) Delete(restaurant models.Restaurant) (models.Restaurant, error) {
	result := r.DB.Delete(&restaurant)

	if result.Error != nil {
		return restaurant, result.Error
	}
	if result.RowsAffected == 0 {
		return restaurant, errors.New("Could not delete data. 0 rows deleted.")
	}
	return restaurant, nil
}