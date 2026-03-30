package booking

import "time"

type Booking struct {
	ID        string
	MovieId   string
	SeatId    string
	UserId    string
	Status    string
	ExpiresAt time.Time
}


type BookingStore interface {
	Book(b Booking) (string, error)
	ListBookings(movieId string) []Booking
	ConfirmBooking(id string) (Booking, error)
}
