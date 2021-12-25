package controllers

import (
	"pedy/controllers/base"
	"pedy/repositories"
	"pedy/services/user"

	"github.com/gin-gonic/gin"
)

func NewUser(c *gin.Context) {
	var dto user.UserDto

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(400, base.Presenter(
			false,
			[]string{err.Error()},
			map[string]interface{}{},
		))
		return
	}

	repo := &repositories.UserRepository{}
	service := user.NewUserService(repo)

	result, err := service.CreateUser(dto)
	if err != nil {
		c.JSON(400, base.Presenter(
			false,
			[]string{err.Error()},
			map[string]interface{}{},
		))
		return
	}

	c.JSON(200, base.Presenter(
		true,
		[]string{},
		result,
	))
}
