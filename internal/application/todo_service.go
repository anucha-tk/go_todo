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

func (s *TodoService) Find(id uint) (*domain.Todo, error) {
	return s.repo.Find(id)
}

func (s *TodoService) Create(todo domain.Todo) (domain.Todo, error) {
	return s.repo.Create(todo)
}

func (s *TodoService) Update(todo domain.Todo, description string, completed bool) (domain.Todo, error) {
	return s.repo.Update(todo, description, completed)
}

func (s *TodoService) Delete(id uint) error {
	return s.repo.Delete(id)
}
