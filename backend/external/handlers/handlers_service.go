package handlers

import (
	"net/http"

	"mafiasu_ws/external/interfaces"
	"mafiasu_ws/external/models"

	"github.com/gin-gonic/gin"
)

type AffiliateHandler struct {
	service interfaces.AffiliateService
}

func NewAffiliateHandler(service interfaces.AffiliateService) *AffiliateHandler {
	return &AffiliateHandler{service}
}

// GetAllAffiliates
func (h *AffiliateHandler) GetAllAffiliates(c *gin.Context) {
	affiliates, err := h.service.GetAllAffiliates(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, affiliates)
}

// GetActiveCars
func (h *AffiliateHandler) GetAvailableCars(c *gin.Context) {
	affiliateID := c.Param("affiliate_id")

	cars, err := h.service.GetAvailableCars(c.Request.Context(), affiliateID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch cars"})
		return
	}

	c.JSON(http.StatusOK, cars)
}

// GetAffiliatesBooking
func (h *AffiliateHandler) GetAffiliateBookingDetail(c *gin.Context) {
	affiliateID := c.Param("affiliate_id")
	bookingID := c.Param("booking_id")

	booking, err := h.service.GetAffiliateBookingDetail(c.Request.Context(), affiliateID, bookingID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, booking)
}

// TrackClick
func (h *AffiliateHandler) TrackClick(c *gin.Context) {
	var req models.TrackClickRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	affiliateID := c.Param("affiliate_id")

	track, err := h.service.TrackClick(c.Request.Context(), affiliateID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to track click"})
		return
	}

	c.JSON(http.StatusOK, models.TrackClickResponse{
		SessionID: track.SessionID,
		Message:   "Click tracked successfully",
	})
}

// AddBookingbyAffiliator
func (h *AffiliateHandler) CreateAffiliateBooking(c *gin.Context) {
	var req models.BookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	affiliateID := c.Param("affiliate_id")

	booking, err := h.service.CreateAffiliateBooking(c.Request.Context(), affiliateID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"booking_id":  booking.BookID,
		"total_price": booking.TotalPrice,
		"message":     "Booking confirmed",
	})
}
