package service

import (
	"context"
	"mafiasu_ws/internal/interfaces"
	"mafiasu_ws/internal/models"
)

type carService struct {
	repo interfaces.CarRepository
}

func NewCarService(repo interfaces.CarRepository) interfaces.CarService {
	return &carService{repo: repo}
}

func (s *carService) GetCarByID(ctx context.Context, id string) (models.Car, error) {
	return s.repo.GetCarByID(ctx, id)
}

func (s *carService) GetAllCars(ctx context.Context) ([]models.Car, error) {
	return s.repo.GetAllCars(ctx)
}

func (s *carService) AddCar(ctx context.Context, car models.CreateCarRequest) (models.Car, error) {
	return s.repo.AddCar(ctx, car)
}

func (s *carService) UpdateCar(ctx context.Context, id string, car models.Car) (models.Car, error) {
	return s.repo.UpdateCar(ctx, id, car)
}

func (s *carService) DeleteCar(ctx context.Context, id string) (models.Car, error) {
	return s.repo.DeleteCar(ctx, id)
}
