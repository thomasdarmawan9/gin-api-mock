package repository

import (
	// "errors"
	"gin-api/models"

	"github.com/stretchr/testify/mock"
)

type CarTest struct {
	Mock mock.Mock
}

//Setiap function harus sama dengan di car repository berikut juga dengan jumlah functionnya
func (repository *CarTest) GetCarById(id int) (*models.Car, error) {
	arguments := repository.Mock.Called(id)

	return arguments.Get(0).(*models.Car), arguments.Error(1)
}

func (repository *CarTest) GetAllCars(data *[]models.Car) (*[]models.Car, error) {
	arguments := repository.Mock.Called(data)

	return arguments.Get(0).(*[]models.Car), arguments.Error(1)
}

func (repository *CarTest) CreateCars(data *models.Car) (*models.Car, error) {
	arguments := repository.Mock.Called(data)

	return arguments.Get(0).(*models.Car), arguments.Error(1)
}
