package models

import "time"

type Car struct {
	CarID             string    `db:"car_id" json:"car_id"`
	Brand             string    `db:"brand" json:"brand"`
	Model             string    `db:"model" json:"model"`
	LicensePlate      string    `db:"license_plate" json:"license_plate"`
	CarType           string    `db:"cartype" json:"cartype"`
	Seat              int       `db:"seat" json:"seat"`
	Doors             int       `db:"doors" json:"doors"`
	GearType          string    `db:"geartype" json:"geartype"`
	FuelType          string    `db:"fueltype" json:"fueltype"`
	RentalPricePerDay float64   `db:"rental_price_per_day" json:"rental_price_per_day"`
	Status            string    `db:"status" json:"status"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time `db:"updated_at" json:"updated_at"`
}

type CreateCarRequest struct {
	Brand             string  `json:"brand" db:"brand"`
	Model             string  `json:"model" db:"model"`
	LicensePlate      string  `json:"license_plate" db:"license_plate"`
	CarType           string  `json:"cartype" db:"cartype"`
	Seat              int     `json:"seat" db:"seat"`
	Doors             int     `json:"doors" db:"doors"`
	GearType          string  `json:"geartype" db:"geartype"`
	FuelType          string  `json:"fueltype" db:"fueltype"`
	RentalPricePerDay float64 `json:"rental_price_per_day" db:"rental_price_per_day"`
	Status            string  `json:"status" db:"status"`
}
