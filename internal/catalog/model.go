package catalog

import "time"

type Location struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `gorm:"size:120;not null" json:"name"` // e.g. "Bengaluru"
	City     string    `gorm:"size:120;not null" json:"city"`
	Theatres []Theatre `json:"theatres,omitempty"`
}

type Theatre struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	Name       string   `gorm:"size:160;not null" json:"name"`
	LocationID uint     `gorm:"index;not null" json:"locationId"`
	Location   Location `json:"location,omitempty"`
	Shows      []Show   `json:"shows,omitempty"`
}

type Movie struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Code        string `gorm:"size:40;uniqueIndex;not null" json:"code"` // keep stable id for frontend/booking
	Title       string `gorm:"size:200;not null" json:"title"`
	Language    string `gorm:"size:40;not null" json:"language"`
	DurationMin int    `gorm:"not null" json:"durationMin"`
	Shows       []Show `json:"shows,omitempty"`
}

type Show struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	MovieID   uint      `gorm:"index;not null" json:"movieId"`
	Movie     Movie     `json:"movie,omitempty"`
	TheatreID uint      `gorm:"index;not null" json:"theatreId"`
	Theatre   Theatre   `json:"theatre,omitempty"`
	StartsAt  time.Time `gorm:"index;not null" json:"startsAt"`
	Format    string    `gorm:"size:20;not null" json:"format"` // 2D/IMAX
	Price     int       `gorm:"not null" json:"price"`          // in smallest currency unit
}

type CatalogStore interface {
	ListLocations() ([]Location, error)
	ListMovies() ([]Movie, error)
	ListShows(movieId uint) ([]Show, error)
	ListTheatres(locationId uint) ([]Theatre, error)
	ListShowsByTheatre(theatreId uint) ([]Show, error)
}
