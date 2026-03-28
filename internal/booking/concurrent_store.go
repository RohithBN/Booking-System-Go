package booking

import (
	"fmt"
	"sync"
)

type ConcurrentStore struct {
	bookings map[string]Booking
	sync.RWMutex
}

func newConcurrentStore() *ConcurrentStore {
	return &ConcurrentStore{
		bookings: map[string]Booking{},
	}
}

func (c *ConcurrentStore) Book(b Booking) error {
	c.Lock()
	defer c.Unlock()
	
	if _, exists := c.bookings[b.SeatId]; exists { 
		return fmt.Errorf("seat %s is already booked", b.SeatId)
	}
	c.bookings[b.SeatId] = b
	return nil
}

func (c *ConcurrentStore) ListBookings(movieId string) []Booking {
	c.RLock()
	defer c.RUnlock()

	var result []Booking
	for _ , b := range c.bookings{
		if b.MovieId == movieId {
			result = append(result, b)
		}
	}
	return result
}