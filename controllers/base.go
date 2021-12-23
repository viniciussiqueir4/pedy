package controllers

import "pedy/repositories"

type BaseController struct {
	RestaurantRepo repositories.RestaurantRepository
}

