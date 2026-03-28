package booking

import "time"

type Booking struct {
	ID      string
	MovieId string
	SeatId  string
	UserId  string
	Status  string
	ExpiresAt time.Time
}


type BookingStore interface{
	Book(b Booking) error
	ListBookings(movieId string) []Booking
}
