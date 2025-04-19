package models

import "time"

type Car struct {
	CarID             string    `db:"car_id"`
	Brand             string    `db:"brand"`
	Model             string    `db:"model"`
	LicensePlate      string    `db:"license_plate"`
	CarType           string    `db:"cartype"`
	Seat              int       `db:"seat"`
	Doors             int       `db:"doors"`
	GearType          string    `db:"geartype"`
	FuelType          string    `db:"fueltype"`
	RentalPricePerDay float64   `db:"rental_price_per_day"`
	Status            string    `db:"status"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}

type CarResponse struct {
	CarID             string  `json:"car_id"`
	Brand             string  `json:"brand"`
	Model             string  `json:"model"`
	LicensePlate      string  `json:"license_plate"`
	CarType           string  `json:"car_type"`
	Seat              int     `json:"seat"`
	Doors             int     `json:"doors"`
	GearType          string  `json:"gear_type"`
	FuelType          string  `json:"fuel_type"`
	RentalPricePerDay float64 `json:"rental_price_per_day"`
	Status            string  `json:"status"`
}

type CreateCarRequest struct {
	Brand             string  `db:"brand"`
	Model             string  `db:"model"`
	LicensePlate      string  `db:"license_plate"`
	CarType           string  `db:"cartype"`
	Seat              int     `db:"seat"`
	Doors             int     `db:"doors"`
	GearType          string  `db:"geartype"`
	FuelType          string  `db:"fueltype"`
	RentalPricePerDay float64 `db:"rental_price_per_day"`
}
