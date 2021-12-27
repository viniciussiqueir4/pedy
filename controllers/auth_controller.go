package controllers

import (
	"pedy/controllers/base"
	"pedy/repositories"
	"pedy/services/auth"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	var dto auth.AuthDto

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
	service := auth.NewAuthService(repo)

	result, err := service.Auth(dto)

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
