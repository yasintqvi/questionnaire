package response

import (
	"github.com/google/uuid"
	"time"
)

type ChoiceResponse struct {
	ID         uuid.UUID `json:"id"`
	QuestionId uuid.UUID `json:"question_id"`
	Value      string    `json:"value"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
