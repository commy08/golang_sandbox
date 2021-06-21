package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/commy08/golang_sandbox.git/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	time.Local = loc

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	connectDB()
}

func main() {
	e := echo.New()
	routes.Routes(e.Group(""))
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func connectDB() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=%s",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
		"Asia%2FBangkok",
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	// Migrate the schema
	// db.AutoMigrate(&models.User{})
	DB = db
}
