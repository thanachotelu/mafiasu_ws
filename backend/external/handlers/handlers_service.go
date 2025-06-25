package handlers

import (
	"net/http"

	"mafiasu_ws/external/interfaces"
	"mafiasu_ws/external/models"
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

type BookingRequest struct {
	BookID     string  `json:"book_id"`
	UserID     string  `json:"user_id"`
	CarID      string  `json:"car_id"`
	SessionID  string  `json:"session_id"`
	TotalPrice float64 `json:"total_price"`
	PickupDate string  `json:"pickup_date,omitempty"`
	ReturnDate string  `json:"return_date,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

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

// @Summary     Get Available Cars
// @Description Get car that are available for booking
// @Tags        Affiliates
// @Accept      json
// @Produce     json
// @Param       affiliate_id path string true "Affiliate ID"
// @Security BearerAuth
// @Success     200 {object} Car
// @Failure     404 {object} ErrorResponse
// @Router      /affiliates/{affiliate_id}/cars [get]
func (h *AffiliateHandler) GetAvailableCars(c *gin.Context) {
	affiliateID := c.Param("affiliate_id")

	cars, err := h.service.GetAvailableCars(c.Request.Context(), affiliateID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch cars"})
		return
	}

	c.JSON(http.StatusOK, cars)
}

// @Summary     Get Affiliate Booking Detail
// @Description Get details of a specific booking made through an affiliate
// @Tags        Affiliates
// @Accept      json
// @Produce     json
// @Param       affiliate_id path string true "Affiliate ID"
// @Param       booking_id path string true "Booking ID"
// @Security BearerAuth
// @Success     200 {object} models.BookingDetailResponse
// @Failure     404 {object} ErrorResponse
// @Router      /affiliates/{affiliate_id}/bookings/{booking_id} [get]
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

// @Summary     Track affiliate click
// @Description Track when an affiliate link is clicked
// @Tags        Affiliates
// @Accept      json
// @Produce     json
// @Param       affiliate_id path string true "Affiliate ID"
// @Param       request body models.TrackClickRequest true "Click tracking info"
// @Security BearerAuth
// @Success     200 {object} models.TrackClickResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /affiliates/{affiliate_id}/track [post]
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

// @Summary     Create affiliate booking
// @Description Create a new booking through affiliate
// @Tags        Affiliates
// @Accept      json
// @Produce     json
// @Param       affiliate_id path string true "Affiliate ID"
// @Param       request body BookingRequest true "Booking details"
// @Security BearerAuth
// @Success     201 {object} models.BookingDetailResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /affiliates/{affiliate_id}/bookings [post]
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
