package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Todo struct {
	gorm.Model
	task   string
	status bool
}

func main() {
	e := echo.New()

	dsn := "host=localhost user=postgres password=password dbname=todos port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("couldnt connect to database.")
	}

	db.AutoMigrate(&Todo{})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})

	e.Logger.Fatal(e.Start(":2024"))
}
