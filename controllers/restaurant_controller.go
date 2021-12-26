package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"pedy/common"
	"pedy/services/restaurant"
)

func (b BaseController) IndexRestaurants(c *gin.Context) {
	c.JSON(200, gin.H{"data": b.RestaurantService.All()})
}

func (b BaseController) GetRestaurant(c *gin.Context) {
	number, err := b.GetIntParam("id", c)
	if err.StatusCode != 0 {
		b.JsonErrors(err, c)
		return
	}
	foundRestaurant, err := b.RestaurantService.Find(number)
	if err.StatusCode != 0 {
		b.JsonErrors(err, c)
		return
	}
	c.JSON(200, gin.H{
		"data": foundRestaurant,
	})
}


func (b BaseController) CreateRestaurant(c *gin.Context) {
	var dto restaurant.RestaurantDTO
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		b.JsonErrors(common.HttpError{
			StatusCode: http.StatusBadRequest,
			Errors:     []error{errors.New("Invalid data provided")},
		}, c)
		return
	}
	create, errs := b.RestaurantService.Create(dto)
	if errs.StatusCode != 0 {
		b.JsonErrors(errs, c)
		return
	}

	c.JSON(http.StatusOK, create)
}

func (b BaseController) UpdateRestaurant(c *gin.Context) {
	id, err := b.GetIntParam("id", c)
	if err.StatusCode != 0 {
		b.JsonErrors(err, c)
		return
	}
	var dto restaurant.RestaurantDTO
	errBind := c.ShouldBindJSON(&dto)
	if errBind != nil {
		b.JsonErrors(common.HttpError{
			StatusCode: http.StatusBadRequest,
			Errors:     []error{errors.New("Invalid data provided")},
		}, c)
		return
	}
	updated, errorsUpdt := b.RestaurantService.Update(dto, id)
	if errorsUpdt.StatusCode != 0 {
		b.JsonErrors(errorsUpdt, c)
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (b BaseController) DeleteRestaurant(c *gin.Context) {
	id, err := b.GetIntParam("id", c)
	if err.StatusCode != 0 {
		b.JsonErrors(err, c)
		return
	}

	find, err := b.RestaurantService.Find(id)
	if err.StatusCode != 0 {
		b.JsonErrors(err, c)
		return
	}
	deletedRestaurant, err := b.RestaurantService.Delete(find)
	if err.StatusCode != 0 {
		b.JsonErrors(err, c)
		return
	}
	c.JSON(http.StatusOK, deletedRestaurant)
}