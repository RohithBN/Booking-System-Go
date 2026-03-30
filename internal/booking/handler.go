package booking

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) HoldSeat() gin.HandlerFunc {
	return func(c *gin.Context) {
		// parse request
		var b Booking
		err := c.ShouldBindJSON(&b)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		if b.MovieId == "" || b.SeatId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "movieId and seatId are required"})
			return
		}
		// call service
		Id,err := h.Service.Book(b)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "seat held successfully", "id": Id})
	}
}

func (h *Handler) ConfirmBooking() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "booking ID is required"})
			return
		}
		b, err := h.Service.ConfirmBooking(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, b)
	}
}

func (h *Handler) ListBookings() gin.HandlerFunc {
	return func(c *gin.Context) {
		movieId := c.Query("movieId")
		if movieId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "movieId query parameter is required"})
			return
		}
		bookings := h.Service.ListBookings(movieId)
		c.JSON(http.StatusOK, bookings)
	}
}
