package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pedy/controllers"
	"pedy/services/restaurant"
)

func ConfigRoutes(router *gin.Engine, db *gorm.DB) *gin.Engine {
	c := controllers.BaseController{
		RestaurantService: restaurant.NewRestaurantService(db),
	}

	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.POST("/", controllers.NewUser)
		}
		restaurants := main.Group("restaurants")
		{
			restaurants.GET("/", c.IndexRestaurants)
			restaurants.GET("/:id", c.GetRestaurant)
			restaurants.POST("/", c.CreateRestaurant)
			restaurants.PUT("/:id", c.UpdateRestaurant)
			restaurants.DELETE("/:id", c.DeleteRestaurant)
		}
	}
	return router
}
