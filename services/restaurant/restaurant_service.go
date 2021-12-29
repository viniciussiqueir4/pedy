package restaurant

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"pedy/common"
	"pedy/interfaces"
	"pedy/models"
	"pedy/repositories"
)

type RestaurantService struct {
	repository interfaces.RestaurantInterface
}

func NewRestaurantService(db *gorm.DB) RestaurantService {
	return RestaurantService{repository: repositories.RestaurantRepository{DB: db}}
}

func (s RestaurantService) All() []models.Restaurant {
	return s.repository.All()
}

func (s RestaurantService) Find(id int) (models.Restaurant, common.HttpError) {
	return s.repository.Find(id)
}

func (s RestaurantService) Create(restaurant RestaurantDTO) (models.Restaurant, common.HttpError) {
	newRestaurant := models.Restaurant{
		Name:   restaurant.Name,
		Cnpj:   restaurant.Cnpj,
		IsOpen: restaurant.IsOpen,
	}

	existentCnpj, httpError := s.repository.FindByCnpj(restaurant.Cnpj)

	if httpError.StatusCode != 0 {
		return existentCnpj, httpError
	}

	if existentCnpj.ID != 0 {
		return newRestaurant, common.HttpError{
			StatusCode: http.StatusBadRequest,
			Errors:     []error{errors.New("Cnpj already exists.")},
		}
	}

	err := newRestaurant.Validate()
	if err != nil {
		return newRestaurant, common.HttpError{
			StatusCode: http.StatusBadRequest,
			Errors:     err,
		}
	}
	return s.repository.Create(newRestaurant)
}

func (s RestaurantService) Delete(restaurant models.Restaurant) (models.Restaurant, common.HttpError) {
	return s.repository.Delete(restaurant)
}

func (s RestaurantService) Update(restaurant RestaurantDTO, id int) (models.Restaurant, common.HttpError) {

	updateRestaurant := models.Restaurant{
		Name:   restaurant.Name,
		Cnpj:   restaurant.Cnpj,
		IsOpen: restaurant.IsOpen,
	}
	existentCnpj, httpError := s.repository.FindByCnpj(restaurant.Cnpj)

	if httpError.StatusCode != 0 {
		return existentCnpj, httpError
	}

	if existentCnpj.ID != 0 {
		return updateRestaurant, common.HttpError{
			StatusCode: http.StatusBadRequest,
			Errors:     []error{errors.New("Cnpj already exists.")},
		}
	}

	err := updateRestaurant.Validate()
	if err != nil {
		return updateRestaurant,  common.HttpError{
			StatusCode: http.StatusBadRequest,
			Errors:     err,
		}
	}

	return s.repository.Update(updateRestaurant, id)
}
