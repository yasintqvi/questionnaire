package domain

import (
	"github.com/google/uuid"
	"time"
)

type UserResponse struct {
	ID           uuid.UUID
	UserId       uuid.UUID
	QuestionId   uuid.UUID
	ResponseType int8
	Choices      []UserResponseChoice
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

func (ur *UserResponse) AddChoice(choice UserResponseChoice) {
	choice.UserResponseId = ur.ID
	ur.Choices = append(ur.Choices, choice)
}
