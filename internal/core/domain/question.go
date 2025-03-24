package domain

import (
	"github.com/google/uuid"
	"time"
)

type Question struct {
	ID              uuid.UUID
	QuestionnaireId uuid.UUID
	Title           string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}
