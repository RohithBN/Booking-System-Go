package booking

import "fmt"

type MemoryStore struct {
	bookings map[string]Booking
}

func newMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}

func (s *MemoryStore) Book(b Booking) error {
	if _, exists := s.bookings[b.SeatId]; exists {
		return fmt.Errorf("seat %s is already booked", b.SeatId)
	}
	s.bookings[b.SeatId] = b
	return nil
}

func (s *MemoryStore) ListBookings(movieId string) []Booking {
	var result []Booking
	for _ , b := range s.bookings{
		if b.MovieId == movieId {
			result = append(result, b)
		}
	}
	return result
}
