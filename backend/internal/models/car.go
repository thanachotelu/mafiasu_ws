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
