package controllers

import (
	"github.com/gin-gonic/gin"
	"pedy/repositories"
)

type BaseController struct {
	RestaurantRepo repositories.RestaurantRepository
}

func (b BaseController) JsonError(statusCode int,err error, c *gin.Context) {
	c.Error(err)
	c.JSON(statusCode, gin.H{"error": err.Error()})
}

