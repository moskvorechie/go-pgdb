package pgdb_test

import (
	"github.com/vmpartner/go-pgdb/v7"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	db, err := pgdb.New(pgdb.Config{
		Host:            "127.0.0.1",
		Port:            30006,
		Name:            "go-pgdb",
		User:            "go-pgdb",
		Pass:            "go-pgdb",
		SslMode:         "disable",
		TimeZone:        "Europe/Moscow",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: time.Hour,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Check query
	err = db.Exec(`SELECT * FROM pg_database`).Error
	if err != nil {
		t.Fatal(err)
	}

	// Create table
	type User struct {
		gorm.Model
		Name     string
		Age      int
		Birthday time.Time
	}
	user := User{Name: "test", Age: 18, Birthday: time.Now()}
	err = db.AutoMigrate(&user)
	if err != nil {
		t.Fatal(err)
	}

	// Truncate
	err = db.Exec(`TRUNCATE TABLE users RESTART IDENTITY CASCADE`).Error
	if err != nil {
		t.Fatal(err)
	}

	// Create user
	err = db.Create(&user).Error
	if err != nil {
		t.Fatal(err)
	}
	if user.ID <= 0 {
		t.Fatal()
	}

	// Get user
	user = User{}
	err = db.First(&user, "name = ? AND age = ?", "test", 18).Error
	if err != nil {
		t.Fatal(err)
	}

	// Drop user
	err = db.Delete(&user).Error
	if err != nil {
		t.Fatal(err)
	}

	// Close check
	bb, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	err = bb.Close()
	if err != nil {
		t.Fatal(err)
	}
	db.Logger = logger.Default.LogMode(logger.Silent)
	err = db.Exec(`SELECT * FROM pg_database`).Error
	if err == nil {
		t.Fatal(err)
	}
}
