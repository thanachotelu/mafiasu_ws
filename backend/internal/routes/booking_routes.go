package routes

import (
	"mafiasu_ws/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterBookingRoutes(r *gin.Engine, bookingHandlers *handler.BookingHandler) {
	bookings := r.Group("/bookings")
	{
		bookings.GET("", bookingHandlers.GetAllBooking)
		bookings.GET("/:id", bookingHandlers.GetBookingByID)
		bookings.POST("", bookingHandlers.AddBooking)
		bookings.PUT("/:id", bookingHandlers.UpdateBooking)
		bookings.DELETE("/:id", bookingHandlers.DeleteBooking)
	}
}
