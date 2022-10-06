package routers

import (
	"gin-api/controllers"
	"gin-api/middleware"
	"gin-api/repository"
	"gin-api/service"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	//Init setiap confignya untuk bisa menjadi turunan dan gampang dibuat mock
	carRepository := repository.NewCarRepository()
	carService := service.NewCarService(carRepository)
	carController := controllers.NewCarController(carService)

	carRoute := router.Group("/cars")
	{
		// middleware for get data by id
		carRoute.GET("/:id", middleware.GetOnevalidation(), carController.GetOneCars)

		// API without middleware
		carRoute.GET("/", carController.GetAllCars)
		carRoute.POST("/cars", carController.CreateCars)
		// router.PATCH("/cars/:id", carController.UpdateCars)
		// router.DELETE("/cars/:id", carController.DeleteCars)
	}

	return router
}
