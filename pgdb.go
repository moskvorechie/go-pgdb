package pgdb

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Config struct {
	Host            string
	Port            int
	Name            string
	User            string
	Pass            string
	Scheme          string
	SslMode         string
	TimeZone        string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func New(c Config) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s search_path=%s TimeZone=%s", c.User, c.Pass, c.Name, c.Host, c.Port, c.SslMode, c.Scheme, c.TimeZone)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	tx, err := db.DB()
	if err != nil {
		return
	}
	tx.SetMaxIdleConns(c.MaxIdleConns)
	tx.SetMaxOpenConns(c.MaxOpenConns)
	tx.SetConnMaxLifetime(c.ConnMaxLifetime)

	return
}
