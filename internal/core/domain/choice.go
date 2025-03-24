package domain

import (
	"github.com/google/uuid"
	"time"
)

type Choice struct {
	ID         uuid.UUID
	QuestionId uuid.UUID
	Value      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
