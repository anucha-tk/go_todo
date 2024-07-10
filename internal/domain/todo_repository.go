package domain

import "gorm.io/gorm"

type TodoRepository interface {
	All(db *gorm.DB) ([]*Todo, error)
}
