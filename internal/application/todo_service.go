package application

import (
	"github.com/anucha-tk/go_todo/internal/domain"
)

type TodoService struct {
	repo domain.TodoRepository
}

func NewTodoService(repo domain.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) FindAll() ([]*domain.Todo, error) {
	return s.repo.FindAll()
}

func (s *TodoService) Delete(id uint) error {
	return s.repo.Delete(id)
}
