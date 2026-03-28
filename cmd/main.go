package main

import (
	"github.com/RohithBN/internal/adapters/redis"
	"github.com/RohithBN/internal/booking"
)


func main(){
	redisClient:= redis.NewRedisClient("localhost:6379")
	store:= booking.NewRedisStore(redisClient)
	service:= booking.NewService(store)
	bookinhHandler:= booking.NewHandler(service)


	  
}