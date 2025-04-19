package interfaces

// import (
// 	"context"
// 	"mafiasu_ws/internal/models"
// )

type CarService interface {
	GetCarByID()
	GetAllCars()
	AddCar()
	UpdateCar()
	DeleteCar()
}
