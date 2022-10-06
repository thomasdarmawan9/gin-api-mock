package controllers

import (
	"gin-api/models"
	"gin-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Kontrak Controllernya
type CarController interface {
	GetAllCars(c *gin.Context)
	GetOneCars(c *gin.Context)
	CreateCars(c *gin.Context)
}

type carControllerImpl struct {
	CarService service.CarService
}

//Inisiasi struct dengan kontrak interface
func NewCarController(newCarService service.CarService) CarController {
	return &carControllerImpl{
		CarService: newCarService,
	}
}

func (controller *carControllerImpl) GetAllCars(c *gin.Context) {
	var cars []models.Car
	res, err := controller.CarService.GetAllCars(&cars)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (controller *carControllerImpl) GetOneCars(c *gin.Context) {
	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter False"})
		return
	}
	res, err := controller.CarService.GetCarById(idint)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data One": res})
}

func (controller *carControllerImpl) CreateCars(c *gin.Context) {
	// Validate input
	var input models.Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create book
	res, err := controller.CarService.CreateCars(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
