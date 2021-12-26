package interfaces

import (
	"pedy/common"
	"pedy/models"
)

type RestaurantInterface interface {
	All() ([]models.Restaurant)
	Find(id int) (models.Restaurant, common.HttpError)
	Delete(restaurant models.Restaurant) (models.Restaurant, common.HttpError)
	Update(restaurant models.Restaurant, idToUpdate int) (models.Restaurant, common.HttpError)
	Create(restaurant models.Restaurant) (models.Restaurant, common.HttpError)
}
