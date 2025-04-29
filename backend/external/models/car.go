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

type CarDetail struct {
	CarID             string  `json:"car_id"`
	Model             string  `json:"model"`
	RentalPricePerDay float64 `json:"rental_price_per_day"`
}
