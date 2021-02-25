# Postgres connect module
Connection module to Postgres database

## Usage example

```golang
package main

import "github.com/moskvorechie/go-pgdb/v7"

func main() {

	// Open connection
	db, err := pgdb.New(pgdb.Config{
		Host:            "127.0.0.1",
		Port:            30006,
		Name:            "go-pgdb",
		User:            "go-pgdb",
		Pass:            "go-pgdb",
		Scheme:          "public",
		SslMode:         "disable",
		TimeZone:        "Europe/Moscow",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: time.Hour,
	})
	if err != nil {
		panic(err)
	}

	// Check query
	err = db.Exec(`SELECT * FROM pg_database`).Error
	if err != nil {
		panic(err)
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
		panic(err)
	}

	// Truncate
	err = db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE").Error
	if err != nil {
		panic(err)
	}

	// Create user
	err = db.Create(&user).Error
	if err != nil {
		panic(err)
	}
	if user.ID <= 0 {
		panic("bad user")
	}

	// Get user
	user = User{}
	err = db.First(&user, "name = ? AND age = ?", "test", 18).Error
	if err != nil {
		panic(err)
	}

	// Drop user
	err = db.Delete(&user).Error
	if err != nil {
		panic(err)
	}
}

```
