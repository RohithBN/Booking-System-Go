package main

import (
	"log"

	"github.com/RohithBN/cmd/server"
	"github.com/RohithBN/internal/adapters/redis"
	"github.com/RohithBN/internal/booking"
	"github.com/RohithBN/internal/catalog"
	"github.com/RohithBN/internal/db"
)

func main() {
	redisClient, err := redis.NewRedisClient("localhost:6379")
	if err != nil {
		panic(err)
	}
	store := booking.NewRedisStore(redisClient)
	service := booking.NewService(store)
	bookingHandler := booking.NewHandler(service)

	cfg := db.Config{
		Hos:      "localhost",
		Port:     "5432",
		User:     "booking",
		Password: "booking",
		Name:     "bookingdb",
		SSLMode:  "disable",
	}
	postgress, err := db.PostgressConnect(cfg)
	if err != nil {
		panic(err)
	}

	if err := postgress.AutoMigrate(&catalog.Location{}, &catalog.Theatre{}, &catalog.Movie{}, &catalog.Show{}); err != nil {
		panic(err)
	}

	postgresStore := catalog.NewPostgressStore(postgress)
	catalogService := catalog.NewService(postgresStore)
	catalogHandler := catalog.NewHandler(catalogService)

	r := server.SetUpRoutes(bookingHandler, catalogHandler)
	log.Println("server starting on :8080")

	r.Run(":8080")

}
