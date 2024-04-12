package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Task   string `json:"task"`
	Status bool   `json:"status"`
}

type TodoPayload struct {
	Task   string `json:"task"`
	Status bool   `json:"status"`
}
