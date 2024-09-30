package domain

import (
	"time"
)

type Todo struct {
	ID          uint
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TodoRepository interface {
	FindAll() ([]*Todo, error)
	Create(todo Todo) (Todo, error)
	Delete(id uint) error
}
