package service

import (
	"errors"
	"fmt"
	"gin-api/models"
	"gin-api/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var carRepository = &repository.CarTest{Mock: mock.Mock{}}
var carServiceTest = NewCarService(carRepository)

func TestCarServiceGetOneProductNotFound(t *testing.T) {
	carRepository.Mock.On("GetCarById", 10).Return(&models.Car{}, errors.New("car not found"))

	product, err := carServiceTest.GetCarById(10)
	if err != nil {
		fmt.Println("error")
	}
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "car not found", err.Error(), "error response has to be 'car not found'")
}

// func TestProductServiceGetOneCar(t *testing.T) {
// 	car := models.Car{
// 		Id:   4,
// 		Merk: "Nissan",
// 		Harga: 7343111111,
// 		Typecars: "K1",
// 	}

// 	carRepository.Mock.On("GetCarById", 4).Return(&models.Car{}, errors.New("car not found"))

// 	result, err := carServiceTest.GetCarById(4)

// 	assert.Nil(t, err)

// 	assert.NotNil(t, result)

// 	assert.Equal(t, car.Id, result.Id, "result has to be '4'")
// 	assert.Equal(t, car.Merk, result.Merk, "result has to be 'Nissan'")
// 	assert.Equal(t, &car, result, "result has to be a product data with id '4'")
// }
