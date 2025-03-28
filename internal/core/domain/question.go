package domain

import (
	"github.com/google/uuid"
	"time"
)

type Question struct {
	ID              uuid.UUID
	QuestionnaireId uuid.UUID
	Title           string
	Choices         []Choice
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

func (q *Question) AddChoice(choice Choice) {

	if choice.ID == uuid.Nil {
		choice.ID = uuid.New()
	}

	choice.QuestionId = q.ID
	q.Choices = append(q.Choices, choice)
}
