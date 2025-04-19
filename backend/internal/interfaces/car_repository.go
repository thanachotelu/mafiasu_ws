package interfaces

// import (
// 	"context"
// 	"mafiasu_ws/internal/models"
// )

type CarRepository interface {
	GetCarByID()
	GetAllCars()
	AddCar()
	UpdateCar()
	DeleteCar()
}
