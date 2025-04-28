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

func NewCarHandler(carService interfaces.CarService) *CarHandler {
	return &CarHandler{carService}
}

func (h *CarHandler) GetCarByID(c *gin.Context) {
	id := c.Param("id")

	car, err := h.carService.GetCarByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, car)
}

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

func (h *CarHandler) DeleteCar(c *gin.Context) {
	id := c.Param("id")

	car, err := h.carService.DeleteCar(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, car)
}

func (h *CarHandler) GetAllCars(c *gin.Context) {
	cars, err := h.carService.GetAllCars(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cars)
}
