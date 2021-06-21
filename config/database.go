package main

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// gorm.Model definition
type Model struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	Model
	Username    string    `grom:"primaryKey; type:text CHARACTER SET utf8 COLLATE uft8_general_ci" json:"username"`
	Password    string    `json:"-"`
	Firstname   string    `gorm:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci" json:"firstname"`
	Lastname    string    `gorm:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci" json:"lastname"`
	Age         int       `json:"age"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

func connectDB() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=%s",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"),
		"Asia%2FBangkok",
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
	DB = db
}
