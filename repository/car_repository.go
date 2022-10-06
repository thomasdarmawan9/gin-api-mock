package repository

import (
	"gin-api/database"
	"gin-api/models"
)

//Kontrak Repositorynya
type CarRepository interface {
	GetCarById(id int) (*models.Car, error)
	GetAllCars(*[]models.Car) (*[]models.Car, error)
	CreateCars(*models.Car) (*models.Car, error)
}

type carRepositoryImpl struct{}

//Inisiasi struct dengan kontrak interface
func NewCarRepository() CarRepository {
	return &carRepositoryImpl{}
}

const (
	queryGetCarById = `SELECT id, merk, harga, typecars FROM cars WHERE id = $1`
	queryGetAllCars = `SELECT id, merk, harga, typecars FROM cars`
	queryCreateCars = `INSERT INTO cars(merk, harga, typecars) VALUES ($1, $2, $3) RETURNING merk, harga, typecars`
)

func (repository *carRepositoryImpl) GetCarById(id int) (*models.Car, error) {
	var db = database.GetDB()

	row := db.Raw(queryGetCarById, id)

	var carOne models.Car

	err := row.Scan(&carOne).Error
	if err != nil {
		return nil, row.Error
	}
	// fmt.Println(&carOne)
	// db.AutoMigrate(&Car{})
	return &carOne, row.Error
}

func (repository *carRepositoryImpl) GetAllCars(*[]models.Car) (*[]models.Car, error) {
	var db = database.GetDB()

	row := db.Raw(queryGetAllCars)

	var carAll []models.Car

	err := row.Scan(&carAll).Error
	if err != nil {
		return nil, row.Error
	}
	return &carAll, row.Error
}

func (repository *carRepositoryImpl) CreateCars(CarReq *models.Car) (*models.Car, error) {
	var db = database.GetDB()

	row := db.Raw(queryCreateCars, CarReq.Merk, CarReq.Harga, CarReq.Typecars)

	var carCreate models.Car

	err := row.Scan(&carCreate).Error
	if err != nil {
		return nil, row.Error
	}

	return &carCreate, row.Error

}
