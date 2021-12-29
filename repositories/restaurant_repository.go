package repositories

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"pedy/common"
	"pedy/models"
)

type RestaurantRepository struct {
	DB *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) RestaurantRepository {
	return RestaurantRepository{DB: db}
}

func (r RestaurantRepository) All() []models.Restaurant {
	var restaurants []models.Restaurant
	r.DB.Find(&restaurants)
	return restaurants
}

func (r RestaurantRepository) Find(id int) (models.Restaurant, common.HttpError) {
	var restaurant models.Restaurant
	result := r.DB.Find(&restaurant, id)
	if result.Error != nil {
		return restaurant, common.HttpError{
			StatusCode: http.StatusInternalServerError,
			Errors:     []error{result.Error},
		}
	}
	if result.RowsAffected == 0 {
		return restaurant, common.HttpError{
			StatusCode: http.StatusNotFound,
			Errors:     []error{errors.New(common.BaseNotFoundText)},
		}
	}
	return restaurant, common.HttpError{}
}

func (r RestaurantRepository) Create(data models.Restaurant) (models.Restaurant, common.HttpError) {
	result := r.DB.Create(&data)
	if result.Error != nil {
		return data, common.HttpError{
			StatusCode: http.StatusInternalServerError,
			Errors:     []error{result.Error},
		}
	}
	if result.RowsAffected == 0 {
		return data, common.HttpError{
			StatusCode: http.StatusBadRequest,
			Errors:     []error{errors.New(common.BaseNoRecordAffected)},
		}
	}
	return data, common.HttpError{}
}

func (r RestaurantRepository) Delete(restaurant models.Restaurant) (models.Restaurant, common.HttpError) {
	result := r.DB.Delete(&restaurant)

	if result.Error != nil {
		return restaurant, common.HttpError{
			StatusCode: http.StatusInternalServerError,
			Errors:     []error{result.Error},
		}
	}
	if result.RowsAffected == 0 {
		return restaurant, common.HttpError{
			StatusCode: http.StatusBadRequest,
			Errors:     []error{errors.New(common.BaseNoRecordAffected)},
		}
	}
	return restaurant, common.HttpError{}
}

func (r RestaurantRepository) Update(newData models.Restaurant, idToUpdate int) (models.Restaurant, common.HttpError) {
	find, err := r.Find(idToUpdate)
	if err.StatusCode != 0 {
		return newData, err
	}
	result := r.DB.Model(&find).Updates(newData)
	if result.Error != nil {
		return newData, common.HttpError{
			StatusCode: http.StatusBadRequest,
			Errors:     []error{result.Error},
		}
	}

	if result.RowsAffected == 0 {
		return newData, common.HttpError{
			StatusCode: http.StatusBadRequest,
			Errors:     []error{errors.New(common.BaseNoRecordAffected)},
		}
	}

	return find, common.HttpError{}
}

func (r RestaurantRepository) FindByCnpj(cnpj string) (models.Restaurant, common.HttpError) {
	var foundRestaurant models.Restaurant
	result := r.DB.Where("cnpj = ?", cnpj).First(&foundRestaurant)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return foundRestaurant, common.HttpError{
			StatusCode: http.StatusInternalServerError,
			Errors:     []error{result.Error},
		}
	}

	return foundRestaurant, common.HttpError{}
}