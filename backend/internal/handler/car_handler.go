package handler

import (
	"mafiasu_ws/internal/interfaces"
	"mafiasu_ws/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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
	Status            string    `db:"status" json:"-"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time `db:"updated_at" json:"updated_at"`
}

type CarHandler struct {
	carService interfaces.CarService
}

type ErrorResponse struct {
	Message string `json:"status"`
}

type SuccessResponse struct {
	Message string `json:"status"`
}

func NewCarHandler(carService interfaces.CarService) *CarHandler {
	return &CarHandler{carService}
}

// @Summary     Get car by ID
// @Description Get car details by ID
// @Tags        Cars
// @Accept      json
// @Produce     json
// @Param       id path string true "Car ID"
// @Success     200 {object} Car
// @Failure     404 {object} ErrorResponse
// @Router      /cars/{id} [get]
func (h *CarHandler) GetCarByID(c *gin.Context) {
	id := c.Param("id")

	car, err := h.carService.GetCarByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, car)
}

// @Summary     Add a car
// @Description Add a new car
// @Tags        Cars
// @Accept      json
// @Produce     json
// @Param       car body models.CreateCarRequest true "Car object"
// @Success     201 {object} Car
// @Failure     500 {object} ErrorResponse
// @Router      /cars [post]
func (h *CarHandler) AddCar(c *gin.Context) {
	var req models.CreateCarRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Status = "active"

	car, err := h.carService.AddCar(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, car)
}

// @Summary     Update a car
// @Description Update car information
// @Tags        Cars
// @Accept      json
// @Produce     json
// @Param       id path string true "Car ID"
// @Param		car body models.CreateCarRequest true "Car object"
// @Success     200 {object} Car
// @Failure     500 {object} ErrorResponse
// @Router      /cars/{id} [put]
func (h *CarHandler) UpdateCar(c *gin.Context) {
	id := c.Param("id")

	// Get existing car first
	existingCar, err := h.carService.GetCarByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "car not found"})
		return
	}

	// Bind the request body to update only provided fields
	if err := c.ShouldBindJSON(&existingCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the car with merged data
	updatedCar, err := h.carService.UpdateCar(c.Request.Context(), id, existingCar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCar)
}

// @Summary     Delete a car
// @Description Deleting a car
// @Tags        Cars
// @Accept      json
// @Produce     json
// @Param       id path string true "Car ID"
// @Success     200 {object} SuccessResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /cars/{id} [delete]
func (h *CarHandler) DeleteCar(c *gin.Context) {
	id := c.Param("id")

	_, err := h.carService.DeleteCar(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Car deleted successfully"})
}

// @Summary     Get all cars
// @Description Get list of all cars
// @Tags        Cars
// @Accept      json
// @Produce     json
// @Success     200 {object} Car
// @Failure     500 {object} ErrorResponse
// @Router      /cars [get]
func (h *CarHandler) GetAllCars(c *gin.Context) {
	cars, err := h.carService.GetAllCars(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cars)
}
