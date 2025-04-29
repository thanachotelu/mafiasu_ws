package handler

import (
	"mafiasu_ws/internal/interfaces"
	"mafiasu_ws/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CarHandler struct {
	carService interfaces.CarService
}

type ErrorResponse struct {
	Message string `json:"message"`
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
// @Success     200 {object} models.Car
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
// @Success     200 {array} models.Car
// @Failure     500 {object} ErrorResponse
// @Router      /cars [post]
func (h *CarHandler) AddCar(c *gin.Context) {
	var req models.CreateCarRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
// @Param       car body models.CreateCarRequest true "Car object"
// @Success     200 {array} models.Car
// @Failure     500 {object} ErrorResponse
// @Router      /cars/{id} [put]
func (h *CarHandler) UpdateCar(c *gin.Context) {
	id := c.Param("id")

	var req models.Car
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car, err := h.carService.UpdateCar(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, car)
}

// @Summary     Delete a car
// @Description Deleting a car
// @Tags        Cars
// @Accept      json
// @Produce     json
// @Param       id path string true "Car ID"
// @Success     200 {array} models.Car
// @Failure 	400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /cars/{id} [delete]
func (h *CarHandler) DeleteCar(c *gin.Context) {
	id := c.Param("id")

	car, err := h.carService.DeleteCar(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, car)
}

// @Summary     Get all cars
// @Description Get list of all cars
// @Tags        Cars
// @Accept      json
// @Produce     json
// @Success     200 {array} models.Car
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
