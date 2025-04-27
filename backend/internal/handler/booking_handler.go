package handler

import (
	"mafiasu_ws/internal/interfaces"
	"mafiasu_ws/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	service interfaces.BookingService
}

func NewBookingHandler(service interfaces.BookingService) *BookingHandler {
	return &BookingHandler{service: service}
}

func (h *BookingHandler) GetBookingByID(c *gin.Context) {
	id := c.Param("id")
	booking, err := h.service.GetBookingByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, booking)
}

func (h *BookingHandler) GetAllBooking(c *gin.Context) {
	bookings, err := h.service.GetAllBooking(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookings)
}

func (h *BookingHandler) AddBooking(c *gin.Context) {
	var req models.BookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	booking, err := h.service.AddBooking(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, booking)
}

func (h *BookingHandler) UpdateBooking(c *gin.Context) {
	id := c.Param("id")

	var req models.Booking
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	booking, err := h.service.UpdateBooking(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, booking)
}

func (h *BookingHandler) DeleteBooking(c *gin.Context) {
	id := c.Param("id")
	booking, err := h.service.DeleteBooking(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, booking)
}
