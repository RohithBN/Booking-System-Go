package server

import (
	"github.com/RohithBN/internal/booking"
	"github.com/RohithBN/internal/catalog"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(h *booking.Handler, c *catalog.Handler) *gin.Engine {
	r := gin.Default()

	r.POST("/hold", h.HoldSeat())
	r.POST("/confirm/:id", h.ConfirmBooking())
	r.GET("/bookings", h.ListBookings())

	r.GET("/locations", c.ListLocations())
	r.GET("/movies", c.ListMovies())
	r.GET("/theatres", c.ListTheatres())
	r.GET("/shows", c.ListShows())

	return r
}
