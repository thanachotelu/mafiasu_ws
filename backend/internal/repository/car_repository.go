package repository

import (
	"context"
	"fmt"
	"mafiasu_ws/internal/interfaces"
	"mafiasu_ws/internal/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type carRepository struct {
	db *pgxpool.Pool
}

func NewCarRepository(db *pgxpool.Pool) interfaces.CarRepository {
	return &carRepository{db: db}
}

func (r *carRepository) GetCarByID(ctx context.Context, id string) (models.Car, error) {
	var car models.Car

	err := r.db.QueryRow(ctx, `
		SELECT car_id, brand, model, license_plate, cartype, seat, doors, geartype, fueltype, rental_price_per_day, status, created_at, updated_at
		FROM cars
		WHERE car_id = $1
	`, id).Scan(
		&car.CarID,
		&car.Brand,
		&car.Model,
		&car.LicensePlate,
		&car.CarType,
		&car.Seat,
		&car.Doors,
		&car.GearType,
		&car.FuelType,
		&car.RentalPricePerDay,
		&car.Status,
		&car.CreatedAt,
		&car.UpdatedAt,
	)

	if err != nil {
		return models.Car{}, fmt.Errorf("car not found: %w", err)
	}

	return car, nil
}

func (r *carRepository) GetAllCars(ctx context.Context) ([]models.Car, error) {
	rows, err := r.db.Query(ctx, `
		SELECT car_id, brand, model, license_plate, cartype, seat, doors, geartype, fueltype, rental_price_per_day, status, created_at, updated_at
		FROM cars
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get cars: %w", err)
	}
	defer rows.Close()

	var cars []models.Car
	for rows.Next() {
		var car models.Car
		err := rows.Scan(
			&car.CarID,
			&car.Brand,
			&car.Model,
			&car.LicensePlate,
			&car.CarType,
			&car.Seat,
			&car.Doors,
			&car.GearType,
			&car.FuelType,
			&car.RentalPricePerDay,
			&car.Status,
			&car.CreatedAt,
			&car.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	return cars, nil
}

func (r *carRepository) AddCar(ctx context.Context, car models.CreateCarRequest) (models.Car, error) {
	var newCar models.Car

	err := r.db.QueryRow(ctx, `
		INSERT INTO cars (brand, model, license_plate, cartype, seat, doors, geartype, fueltype, rental_price_per_day, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, 'active', NOW(), NOW())
		RETURNING car_id, brand, model, license_plate, cartype, seat, doors, geartype, fueltype, rental_price_per_day, status, created_at, updated_at
	`,
		car.Brand, car.Model, car.LicensePlate, car.CarType, car.Seat, car.Doors, car.GearType, car.FuelType, car.RentalPricePerDay,
	).Scan(
		&newCar.CarID,
		&newCar.Brand,
		&newCar.Model,
		&newCar.LicensePlate,
		&newCar.CarType,
		&newCar.Seat,
		&newCar.Doors,
		&newCar.GearType,
		&newCar.FuelType,
		&newCar.RentalPricePerDay,
		&newCar.Status,
		&newCar.CreatedAt,
		&newCar.UpdatedAt,
	)

	if err != nil {
		return models.Car{}, fmt.Errorf("failed to add car: %w", err)
	}

	return newCar, nil
}

func (r *carRepository) UpdateCar(ctx context.Context, id string, car models.Car) (models.Car, error) {
	var updatedCar models.Car

	err := r.db.QueryRow(ctx, `
		UPDATE cars
		SET brand=$1, model=$2, license_plate=$3, cartype=$4, seat=$5, doors=$6, geartype=$7, fueltype=$8, rental_price_per_day=$9, status=$10, updated_at=NOW()
		WHERE car_id=$11
		RETURNING car_id, brand, model, license_plate, cartype, seat, doors, geartype, fueltype, rental_price_per_day, status, created_at, updated_at
	`,
		car.Brand, car.Model, car.LicensePlate, car.CarType, car.Seat, car.Doors, car.GearType, car.FuelType, car.RentalPricePerDay, car.Status, id,
	).Scan(
		&updatedCar.CarID,
		&updatedCar.Brand,
		&updatedCar.Model,
		&updatedCar.LicensePlate,
		&updatedCar.CarType,
		&updatedCar.Seat,
		&updatedCar.Doors,
		&updatedCar.GearType,
		&updatedCar.FuelType,
		&updatedCar.RentalPricePerDay,
		&updatedCar.Status,
		&updatedCar.CreatedAt,
		&updatedCar.UpdatedAt,
	)

	if err != nil {
		return models.Car{}, fmt.Errorf("failed to update car: %w", err)
	}

	return updatedCar, nil
}

func (r *carRepository) DeleteCar(ctx context.Context, id string) (models.Car, error) {
	var deletedCar models.Car

	err := r.db.QueryRow(ctx, `
		DELETE FROM cars
		WHERE car_id = $1
		RETURNING car_id, brand, model, license_plate, cartype, seat, doors, geartype, fueltype, rental_price_per_day, status, created_at, updated_at
	`, id).Scan(
		&deletedCar.CarID,
		&deletedCar.Brand,
		&deletedCar.Model,
		&deletedCar.LicensePlate,
		&deletedCar.CarType,
		&deletedCar.Seat,
		&deletedCar.Doors,
		&deletedCar.GearType,
		&deletedCar.FuelType,
		&deletedCar.RentalPricePerDay,
		&deletedCar.Status,
		&deletedCar.CreatedAt,
		&deletedCar.UpdatedAt,
	)

	if err != nil {
		return models.Car{}, fmt.Errorf("failed to delete car: %w", err)
	}

	return deletedCar, nil
}
