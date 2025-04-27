package interfaces

import (
	"context"
	"mafiasu_ws/internal/models"
)

type CarService interface {
	GetCarByID(ctx context.Context, id string) (models.Car, error)
	GetAllCars(ctx context.Context) ([]models.Car, error)
	AddCar(ctx context.Context, car models.CreateCarRequest) (models.Car, error)
	UpdateCar(ctx context.Context, id string, car models.Car) (models.Car, error)
	DeleteCar(ctx context.Context, id string) (models.Car, error)
}
