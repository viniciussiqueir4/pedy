package routes

import (
	"pedy/controllers"
	"pedy/repositories"

	"pedy/server/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigRoutes(router *gin.Engine, db *gorm.DB) *gin.Engine {
	c := controllers.BaseController{
		RestaurantRepo: repositories.NewRestaurantRepository(db),
	}

	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.POST("/", controllers.NewUser)
			users.GET("/", middlewares.Auth(), controllers.GetUserById)
		}
		auth := main.Group("auth")
		{
			auth.POST("/", controllers.Auth)
		}
		restaurants := main.Group("restaurants")
		{
			restaurants.GET("/", c.IndexRestaurants)
			restaurants.GET("/:id", c.GetRestaurant)
			restaurants.POST("/", c.CreateRestaurant)
			restaurants.DELETE("/:id", c.DeleteRestaurant)
		}
	}
	return router
}
