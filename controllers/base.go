package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"pedy/common"
	"pedy/services/restaurant"
	"strconv"
)




type BaseController struct {
	RestaurantService restaurant.RestaurantService
}

//Accepts an array of error
func (b BaseController) JsonErrors(errs common.HttpError, c *gin.Context) {
	strErrors := make([]string, len(errs.Errors))
	for i, err := range errs.Errors {
		strErrors[i] = err.Error()
		c.Error(err)
	}
	c.JSON(errs.StatusCode, gin.H{"errors": strErrors})
}

//gets an integer parameter from the request
func (b BaseController) GetIntParam(paramName string,c *gin.Context) (int, common.HttpError) {
	id := c.Param(paramName)
	errParameter := errors.New("Invalid parameter.")
	if id == "" {
		return 0, common.HttpError{
			StatusCode: http.StatusBadRequest,
			Errors:     []error{errParameter},
		}
	}
	number, err := strconv.Atoi(id)
	if err != nil {
		return 0, common.HttpError{
			StatusCode: http.StatusBadRequest,
			Errors:     []error{errParameter},
		}
	}
	return number, common.HttpError{}
}