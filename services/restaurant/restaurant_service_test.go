package restaurant

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"net/http"
	"pedy/common"
	"pedy/models"
	"testing"
	"time"
)

type RestaurantRepositoryInterfaceMock struct {
	mock.Mock
}

func (m RestaurantRepositoryInterfaceMock) All() []models.Restaurant {
	args := m.Called()
	return args.Get(0).([]models.Restaurant)
}

func (m RestaurantRepositoryInterfaceMock) FindByCnpj(cnpj string) (models.Restaurant, common.HttpError) {
	args := m.Called(cnpj)
	return args.Get(0).(models.Restaurant), args.Get(1).(common.HttpError)
}

func (m RestaurantRepositoryInterfaceMock) Delete(restaurant models.Restaurant) (models.Restaurant, common.HttpError) {
	args := m.Called(restaurant)
	return args.Get(0).(models.Restaurant), args.Get(1).(common.HttpError)
}

func (m RestaurantRepositoryInterfaceMock) Update(restaurant models.Restaurant, idToUpdate int) (models.Restaurant, common.HttpError) {
	args := m.Called(restaurant, idToUpdate)
	return args.Get(0).(models.Restaurant), args.Get(1).(common.HttpError)
}

func (m RestaurantRepositoryInterfaceMock) Create(restaurant models.Restaurant) (models.Restaurant, common.HttpError) {
	args := m.Called(restaurant)
	return args.Get(0).(models.Restaurant), args.Get(1).(common.HttpError)
}

func (m RestaurantRepositoryInterfaceMock) Find(id int) (models.Restaurant, common.HttpError) {
	args := m.Called(id)
	return args.Get(0).(models.Restaurant), args.Get(1).(common.HttpError)
}

func TestItFindsRestaurant(t *testing.T) {
	repository := new(RestaurantRepositoryInterfaceMock)
	service := RestaurantService{repository: repository}

	repository.On("Find", 1).Return(models.Restaurant{
		ID:        1,
		Name:      "Teste 123",
		Cnpj:      "",
		IsOpen:    false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}, common.HttpError{})

	find, _ := service.Find(1)
	repository.AssertExpectations(t)
	assert.Equal(t, "Teste 123", find.Name, "They should be equal")
}

func TestItFindsRestaurants(t *testing.T) {
	repository := new(RestaurantRepositoryInterfaceMock)
	service := RestaurantService{repository: repository}

	repository.On("All").Return([]models.Restaurant{
		{
			ID:        1,
			Name:      "teste 1",
			Cnpj:      "",
			IsOpen:    false,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		{
			ID:        2,
			Name:      "teste 2",
			Cnpj:      "",
			IsOpen:    false,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
	},
	)

	find := service.All()
	repository.AssertExpectations(t)
	assert.Equal(t, 2, len(find), "They should be equal")
}

func TestItDeletesRestaurant(t *testing.T) {
	repository := new(RestaurantRepositoryInterfaceMock)
	service := RestaurantService{repository: repository}
	toDelete := models.Restaurant{
		Name:   "teste",
		Cnpj:   "123",
		IsOpen: false,
	}
	repository.On("Delete", toDelete).Return(toDelete, common.HttpError{})

	find, err := service.Delete(toDelete)
	repository.AssertExpectations(t)
	assert.Equal(t, "teste", find.Name, "They should be equal")
	assert.Empty(t, err, "A error should not be thrown.")
}

func TestItCantCreateRestaurantIfCnpjExists(t *testing.T) {
	repository := new(RestaurantRepositoryInterfaceMock)
	service := RestaurantService{repository: repository}
	toCreate := RestaurantDTO{
		Name:   "teste",
		Cnpj:   "123",
		IsOpen: true,
	}

	foundRestaurant := models.Restaurant{
		ID:     1,
		Name:   "teste",
		Cnpj:   "123",
		IsOpen: true,
	}
	repository.On("FindByCnpj", foundRestaurant.Cnpj).Return(foundRestaurant, common.HttpError{})

	_, err := service.Create(toCreate)
	repository.AssertExpectations(t)
	assert.Equal(t, http.StatusBadRequest, err.StatusCode, "They should be equal")
	assert.Len(t, err.Errors, 1, "Should have one error.")
}

func TestItCanCreateRestaurant(t *testing.T) {
	repository := new(RestaurantRepositoryInterfaceMock)
	service := RestaurantService{repository: repository}
	toCreate := RestaurantDTO{
		Name:   "teste",
		Cnpj:   "35784409000120",
		IsOpen: true,
	}

	repository.On("FindByCnpj", "35784409000120").Return(models.Restaurant{}, common.HttpError{})
	restaurant := models.Restaurant{
		Name:   toCreate.Name,
		Cnpj:   toCreate.Cnpj,
		IsOpen: toCreate.IsOpen,
	}
	repository.On("Create", restaurant).Return(restaurant, common.HttpError{})
	created, err := service.Create(toCreate)
	repository.AssertExpectations(t)
	assert.Equal(t, toCreate.Name, created.Name, "They should be equal")
	assert.Empty(t, err, "Error should not be thrown")
}

func TestItValidatesCnpj(t *testing.T) {
	repository := new(RestaurantRepositoryInterfaceMock)
	service := RestaurantService{repository: repository}
	toCreate := RestaurantDTO{
		Name:   "teste",
		Cnpj:   "123",
		IsOpen: true,
	}

	repository.On("FindByCnpj", "123").Return(models.Restaurant{}, common.HttpError{})

	created, err := service.Create(toCreate)
	repository.AssertExpectations(t)
	assert.Equal(t, toCreate.Name, created.Name, "They should be equal")
	assert.NotEmpty(t, err, "Error should not be empty")
	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
}

func TestItCanUpdateRestaurant(t *testing.T) {
	repository := new(RestaurantRepositoryInterfaceMock)
	service := RestaurantService{repository: repository}
	toUpdate := RestaurantDTO{
		Name:   "teste",
		Cnpj:   "35784409000120",
		IsOpen: true,
	}

	repository.On("FindByCnpj", "35784409000120").Return(models.Restaurant{}, common.HttpError{})
	restaurant := models.Restaurant{
		Name:   toUpdate.Name,
		Cnpj:   toUpdate.Cnpj,
		IsOpen: toUpdate.IsOpen,
	}
	repository.On("Update", restaurant, 1).Return(
		restaurant,
		common.HttpError{},
	)
	created, err := service.Update(toUpdate, 1)
	repository.AssertExpectations(t)
	assert.Equal(t, toUpdate.Name, created.Name, "They should be equal")
	assert.Empty(t, err, "Error should not be thrown")
}
