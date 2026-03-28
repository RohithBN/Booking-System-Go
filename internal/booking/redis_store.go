package booking

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const defaultHoldTime = 2 * time.Minute

type RedisStore struct {
	rdb *redis.Client
}

func NewRedisStore(rdb *redis.Client) *RedisStore {
	return &RedisStore{
		rdb: rdb,
	}
}

func sessionKey(id string) string {
	return fmt.Sprintf("session:%s", id)
}

func (s *RedisStore) Book(b Booking) error {
	res, err := s.hold(b)
	if err != nil {
		return err
	}
	fmt.Println("Seat Booked: ", res)
	// once booking is confirmed we can update the status to "booked" and remove the session key
	res, err = s.confirmBooking(res.ID)
	return nil
}

func (s *RedisStore) ListBookings(movieId string) []Booking { return nil }

func (s *RedisStore) hold(b Booking) (Booking, error) {
	id := uuid.New().String()
	now := time.Now()
	ctx := context.Background()
	key := fmt.Sprintf("booking:%s:%s", b.MovieId, b.SeatId)
	b.ID = id
	b.Status = "held"
	b.ExpiresAt = now.Add(defaultHoldTime)

	val, err := json.Marshal(b)
	if err != nil {
		return Booking{}, err
	}
	res := s.rdb.SetArgs(ctx, key, val, redis.SetArgs{
		Mode: "NX",
		TTL:  defaultHoldTime,
	})
	ok := res.Val() == "OK"
	if !ok {
		return Booking{}, fmt.Errorf("seat %s is already booked", b.SeatId)
	}

	s.rdb.Set(ctx, sessionKey(id), key, defaultHoldTime)

	return Booking{
		ID:        id,
		MovieId:   b.MovieId,
		SeatId:    b.SeatId,
		UserId:    b.UserId,
		Status:    "held",
		ExpiresAt: now.Add(defaultHoldTime),
	}, nil
}

func (s *RedisStore) confirmBooking(id string) (Booking, error) {
	ctx := context.Background()
	sessionKey := sessionKey(id)
	key, err := s.rdb.Get(ctx, sessionKey).Result()
	if err != nil {
		return Booking{}, fmt.Errorf("invalid session")
	}
	val, err := s.rdb.Get(ctx, key).Result()
	if err != nil {
		return Booking{}, fmt.Errorf("booking expired")
	}
	var b Booking
	err = json.Unmarshal([]byte(val), &b)
	if err != nil {
		return Booking{}, fmt.Errorf("failed to parse booking data")
	}
	b.Status = "booked"
	valBytes, err := json.Marshal(b)
	if err != nil {
		return Booking{}, fmt.Errorf("failed to serialize booking data")
	}
	s.rdb.Set(ctx, key, valBytes, 0) // update the booking status to "booked"
	s.rdb.Del(ctx, sessionKey)       // remove the session key
	return b, nil
}
