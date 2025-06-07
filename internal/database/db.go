package db

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=root dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&User{})
}
