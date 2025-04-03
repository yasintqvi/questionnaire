package domain

import (
	"github.com/google/uuid"
	"time"
)

type Questionnaire struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Status      bool      `db:"status"`
	StartTime   time.Time `db:"start_time"`
	EndTime     time.Time `db:"end_time"`
	Questions   []Question
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}

func (q *Questionnaire) AddQuestion(question Question) {

	if question.ID == uuid.Nil {
		question.ID = uuid.New()
	}

	question.QuestionnaireId = q.ID
	q.Questions = append(q.Questions, question)
}
