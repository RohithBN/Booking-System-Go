package catalog

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ListLocations() gin.HandlerFunc {
	return func(c *gin.Context) {

		locations, err := h.service.ListLocations()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, locations)
	}
}

func (h *Handler) ListMovies() gin.HandlerFunc {
	return func(c *gin.Context) {

		movies, err := h.service.ListMovies()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, movies)
	}
}

func (h *Handler) ListShows() gin.HandlerFunc {
	return func(c *gin.Context) {
		movieIdStr := c.Query("movieId")
		theatreIdStr := c.Query("theatreId")

		if movieIdStr == "" && theatreIdStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "movieId or theatreId is required"})
			return
		}

		if movieIdStr != "" {
			movieId, err := parseUint(movieIdStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid movieId"})
				return
			}
			shows, err := h.service.ListShows(movieId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, shows)
			return
		}

		theatreId, err := parseUint(theatreIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid theatreId"})
			return
		}
		shows, err := h.service.ListShowsByTheatre(theatreId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, shows)
	}
}

func parseUint(s string) (uint, error) {
	var num uint
	_, err := fmt.Sscanf(s, "%d", &num)
	return num, err
}

func (h *Handler) ListTheatres() gin.HandlerFunc {
	return func(c *gin.Context) {
		locationIdStr := c.Query("locationId")
		if locationIdStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "locationId is required"})
			return
		}
		locationId, err := parseUint(locationIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid locationId"})
			return
		}
		theatres, err := h.service.ListTheatres(locationId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, theatres)
	}
}
