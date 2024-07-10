package home

import (
	"context"

	"github.com/anucha-tk/go_todo/internal/database"
	"github.com/anucha-tk/go_todo/internal/domain"
)

type (
	Service interface {
		List(ctx context.Context) ([]*domain.Todo, error)
	}
	service struct {
		todos domain.TodoRepository
	}
)

func NewService() Service {
	todoRepo := &domain.Todos{}
	return &service{
		todos: todoRepo,
	}
}

func (s service) List(context.Context) ([]*domain.Todo, error) {
	db := database.GetDB()
	todos, err := s.todos.All(db)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
