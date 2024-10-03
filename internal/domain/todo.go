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
	Find(id uint) (*Todo, error)
	Create(todo Todo) (Todo, error)
	Update(todo Todo, description string, complated bool) (Todo, error)
	Delete(id uint) error
}
