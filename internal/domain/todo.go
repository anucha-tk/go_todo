package domain

import (
	"time"
)

type Todo struct {
	ID          int
	Description string
	Completed   bool
	CreatedAt   time.Time
}
