package service

import (
	"gin-api/models"
	"gin-api/repository"
)

//Kontrak Servicenya
type CarService interface {
	GetCarById(id int) (*models.Car, error)
	GetAllCars(mc *[]models.Car) (*[]models.Car, error)
	CreateCars(carReq *models.Car) (*models.Car, error)
}

type carServiceImpl struct {
	CarRepository repository.CarRepository
}

//Inisiasi struct dengan kontrak interface
func NewCarService(newCarReposiory repository.CarRepository) CarService {
	return &carServiceImpl{
		CarRepository: newCarReposiory,
	}
}

func (service *carServiceImpl) GetCarById(id int) (*models.Car, error) {
	car, err := service.CarRepository.GetCarById(id)
	// fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return car, err
}

func (service *carServiceImpl) GetAllCars(mc *[]models.Car) (*[]models.Car, error) {
	car, err := service.CarRepository.GetAllCars(mc)

	if err != nil {
		return nil, err
	}

	return car, err
}

func (service *carServiceImpl) CreateCars(carReq *models.Car) (*models.Car, error) {
	car, err := service.CarRepository.CreateCars(carReq)

	if err != nil {
		return nil, err
	}

	return car, err
}
