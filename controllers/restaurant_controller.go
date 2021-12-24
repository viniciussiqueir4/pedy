package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"pedy/models"
	"strconv"
	"time"
)

func (b BaseController) IndexRestaurants(c *gin.Context) {
	c.JSON(200, gin.H{"data": b.RestaurantRepo.All()})
}

func (b BaseController) GetRestaurant(c *gin.Context) {
	id := c.Param("id")
	errParameter := errors.New("Invalid parameter.")
	if id == ""{
		b.JsonError(http.StatusBadRequest, errParameter, c)
		return
	}
	number, err := strconv.Atoi(id)
	if err != nil {
		b.JsonError(http.StatusBadRequest, errParameter, c)
		return
	}
	restaurant, err := b.RestaurantRepo.Find(number)
	if err != nil{
		b.JsonError(http.StatusNotFound, err, c)
		return
	}
	c.JSON(200, gin.H{
		"data": restaurant,
	})
}
//TODO: finish create with validations
func (b BaseController) CreateRestaurant(c *gin.Context) {
	create, err := b.RestaurantRepo.Create(models.Restaurant{
		Name:      "xd",
		Cnpj:      "123456789123456798",
		IsOpen:    false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	})
	if err != nil {
		b.JsonError(http.StatusBadRequest, err, c)
		return
	}
	c.JSON(http.StatusOK, create)
}

func (b BaseController) DeleteRestaurant(c *gin.Context) {
	id := c.Param("id")
	errParameter := errors.New("Invalid parameter.")
	if id == ""{
		b.JsonError(http.StatusBadRequest, errParameter, c)
		return
	}
	number, err := strconv.Atoi(id)
	if err != nil {
		b.JsonError(http.StatusBadRequest, errParameter, c)
		return
	}
	find, err := b.RestaurantRepo.Find(number)
	if err != nil {
		b.JsonError(http.StatusNotFound, err, c)
		return
	}
	restaurant, err := b.RestaurantRepo.Delete(find)
	if err != nil {
		b.JsonError(http.StatusBadRequest, err, c)
		return
	}
	c.JSON(http.StatusOK, restaurant)
}