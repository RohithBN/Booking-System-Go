package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Hos      string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func PostgressConnect(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Hos, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
