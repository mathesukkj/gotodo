package main

import (
	"log"

	"github.com/labstack/echo/v4"

	"gotodo/db"
	"gotodo/handlers"
)

var e *echo.Echo

func main() {
	e = echo.New()

	if err := db.Init(); err != nil {
		log.Fatal("couldnt connect to the database!")
	}

	e.GET("/todo", handlers.GetTodos)
	e.GET("/todo/:id", handlers.GetTodo)
	e.POST("/todo", handlers.CreateTodo)

	e.Logger.Fatal(e.Start(":2023"))
}
