package main

import (
	"log"
	"time"

	"github.com/RohithBN/internal/catalog"
	"github.com/RohithBN/internal/db"
	"gorm.io/gorm"
)

func main() {
	cfg := db.Config{
		Hos:      "localhost",
		Port:     "5432",
		User:     "booking",
		Password: "booking",
		Name:     "bookingdb",
		SSLMode:  "disable",
	}

	postgresDB, err := db.PostgressConnect(cfg)
	if err != nil {
		log.Fatalf("failed to connect postgres: %v", err)
	}

	if err := postgresDB.AutoMigrate(&catalog.Location{}, &catalog.Theatre{}, &catalog.Movie{}, &catalog.Show{}); err != nil {
		log.Fatalf("failed to migrate tables: %v", err)
	}

	if err := postgresDB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&catalog.Show{}).Error; err != nil {
		log.Fatalf("failed to clear shows: %v", err)
	}
	if err := postgresDB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&catalog.Theatre{}).Error; err != nil {
		log.Fatalf("failed to clear theatres: %v", err)
	}
	if err := postgresDB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&catalog.Movie{}).Error; err != nil {
		log.Fatalf("failed to clear movies: %v", err)
	}
	if err := postgresDB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&catalog.Location{}).Error; err != nil {
		log.Fatalf("failed to clear locations: %v", err)
	}

	locations := []catalog.Location{
		{Name: "Mumbai", City: "Mumbai"},
		{Name: "Bengaluru", City: "Bengaluru"},
		{Name: "Hyderabad", City: "Hyderabad"},
	}
	if err := postgresDB.Create(&locations).Error; err != nil {
		log.Fatalf("failed to seed locations: %v", err)
	}

	theatres := []catalog.Theatre{
		{Name: "PVR Phoenix", LocationID: locations[0].ID},
		{Name: "INOX R-City", LocationID: locations[0].ID},
		{Name: "PVR Orion Mall", LocationID: locations[1].ID},
		{Name: "Cinepolis Forum", LocationID: locations[1].ID},
		{Name: "AMB Cinemas", LocationID: locations[2].ID},
	}
	if err := postgresDB.Create(&theatres).Error; err != nil {
		log.Fatalf("failed to seed theatres: %v", err)
	}

	movies := []catalog.Movie{
		{Code: "INT-RE", Title: "Interstellar Re-Release", Language: "English", DurationMin: 169},
		{Code: "DUNE-2", Title: "Dune: Part Two", Language: "English", DurationMin: 166},
		{Code: "KALKI", Title: "Kalki 2898 AD", Language: "Hindi", DurationMin: 181},
	}
	if err := postgresDB.Create(&movies).Error; err != nil {
		log.Fatalf("failed to seed movies: %v", err)
	}

	now := time.Now()
	showTimes := []time.Time{
		time.Date(now.Year(), now.Month(), now.Day(), 10, 30, 0, 0, now.Location()),
		time.Date(now.Year(), now.Month(), now.Day(), 14, 0, 0, 0, now.Location()),
		time.Date(now.Year(), now.Month(), now.Day(), 18, 30, 0, 0, now.Location()),
		time.Date(now.Year(), now.Month(), now.Day(), 21, 45, 0, 0, now.Location()),
	}

	shows := []catalog.Show{
		{MovieID: movies[0].ID, TheatreID: theatres[0].ID, StartsAt: showTimes[0], Format: "IMAX", Price: 42000},
		{MovieID: movies[1].ID, TheatreID: theatres[0].ID, StartsAt: showTimes[2], Format: "4K", Price: 35000},
		{MovieID: movies[2].ID, TheatreID: theatres[0].ID, StartsAt: showTimes[3], Format: "2D", Price: 25000},
		{MovieID: movies[1].ID, TheatreID: theatres[1].ID, StartsAt: showTimes[1], Format: "4K", Price: 32000},
		{MovieID: movies[0].ID, TheatreID: theatres[2].ID, StartsAt: showTimes[2], Format: "IMAX", Price: 39000},
		{MovieID: movies[2].ID, TheatreID: theatres[3].ID, StartsAt: showTimes[1], Format: "2D", Price: 24000},
		{MovieID: movies[1].ID, TheatreID: theatres[4].ID, StartsAt: showTimes[3], Format: "4K", Price: 36000},
	}
	if err := postgresDB.Create(&shows).Error; err != nil {
		log.Fatalf("failed to seed shows: %v", err)
	}

	log.Println("catalog seed complete")
}
