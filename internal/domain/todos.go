package domain

import (
	"gorm.io/gorm"
)

type Todos struct{}

func (t *Todos) All(db *gorm.DB) ([]*Todo, error) {
	var todos []*Todo
	err := db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}
