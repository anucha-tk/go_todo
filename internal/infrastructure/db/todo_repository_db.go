package db

import (
	"github.com/anucha-tk/go_todo/internal/domain"
	"gorm.io/gorm"
)

type TodoRepositoryGORM struct {
	db *gorm.DB
}

func NewTodoRepositoryGORM(db *gorm.DB) *TodoRepositoryGORM {
	return &TodoRepositoryGORM{db: db}
}

func (r *TodoRepositoryGORM) Save(todo *domain.Todo) error {
	return r.db.Create(todo).Error
}

func (r *TodoRepositoryGORM) FindAll() ([]*domain.Todo, error) {
	var todos []*domain.Todo
	err := r.db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepositoryGORM) Find(id uint) (*domain.Todo, error) {
	var todo *domain.Todo
	err := r.db.Where("id = ?", id).First(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoRepositoryGORM) Update(todo domain.Todo, description string, completed bool) (domain.Todo, error) {
	err := r.db.Model(&todo).Updates(map[string]interface{}{"description": description, "completed": completed}).Error
	if err != nil {
		return domain.Todo{}, err
	}
	return todo, nil
}

func (r *TodoRepositoryGORM) Create(todo domain.Todo) (domain.Todo, error) {
	err := r.db.Create(&todo).Error
	if err != nil {
		return domain.Todo{}, err
	}
	return todo, nil
}

func (r *TodoRepositoryGORM) Delete(id uint) error {
	return r.db.Delete(&domain.Todo{}, id).Error
}
