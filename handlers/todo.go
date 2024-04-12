package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"gotodo/db"
	"gotodo/models"
)

func GetTodos(c echo.Context) error {
	todos := []models.Todo{}

	db.Gorm.Find(&todos)

	return c.JSON(http.StatusOK, todos)
}

func GetTodo(c echo.Context) error {
	todo := models.Todo{}

	result := db.Gorm.First(&todo, c.Param("id"))
	if result.Error != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, todo)
}

func CreateTodo(c echo.Context) error {
	todo := models.TodoPayload{}
	err := json.NewDecoder(c.Request().Body).Decode(&todo)
	if err != nil {
		log.Fatal("error while decoding json!")
	}

	db.Gorm.Create(&models.Todo{Task: todo.Task, Status: false})

	return c.NoContent(http.StatusCreated)
}
