package domain

import (
	"time"
)

type Todo struct {
	ID          uint
	Description string
	Completed   bool
	CreatedAt   time.Time
}

type TodoRepository interface {
	FindAll() ([]*Todo, error)
	Delete(id uint) error
}

func NewTodo(description string) *Todo {
	return &Todo{
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
}

// TODO: make Completed func
