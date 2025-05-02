package routes

import (
	"mafiasu_ws/external/handlers"
	"mafiasu_ws/external/interfaces"

	"github.com/gin-gonic/gin"
)

func RegisterAffiliateRoutes(router *gin.Engine, repo interfaces.AffiliateService) {
	handler := handlers.NewAffiliateHandler(repo)
	v1 := router.Group("/api/v1")
	{
		affiliates := v1.Group("/affiliates")
		{
			affiliates.GET("", handler.GetAllAffiliates)                                             //
			affiliates.GET("/:affiliate_id/cars", handler.GetAvailableCars)                          //
			affiliates.GET("/:affiliate_id/bookings/:booking_id", handler.GetAffiliateBookingDetail) //
			affiliates.POST("/:affiliate_id/track-click", handler.TrackClick)                        //
			affiliates.POST("/:affiliate_id/bookings", handler.CreateAffiliateBooking)
		}
	}
}
