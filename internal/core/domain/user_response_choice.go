package domain

import (
	"github.com/google/uuid"
	"time"
)

type UserResponseChoice struct {
	ID             uuid.UUID
	ChoiceId       uuid.UUID
	UserResponseId uuid.UUID
	CreatedAt      time.Time
	DeletedAt      *time.Time
}
