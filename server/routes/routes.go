package routes

import (
	"pedy/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.POST("/", controllers.NewUser)
		}
	}
	return router
}
