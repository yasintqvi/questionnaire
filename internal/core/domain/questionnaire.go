package domain

import (
	"github.com/google/uuid"
	"time"
)

type Questionnaire struct {
	ID          uuid.UUID
	Title       string
	Description string
	Status      bool
	StartTime   time.Time
	EndTime     time.Time
	Questions   []Question
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func (q *Questionnaire) AddQuestion(question Question) {

	if question.ID == uuid.Nil {
		question.ID = uuid.New()
	}

	question.QuestionnaireId = q.ID
	q.Questions = append(q.Questions, question)
}
