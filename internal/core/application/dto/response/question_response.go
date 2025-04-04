package response

import (
	"github.com/google/uuid"
	"time"
)

type QuestionResponse struct {
	ID              uuid.UUID `json:"id"`
	QuestionnaireId uuid.UUID `json:"questionnaire_id"`
	Title           string    `json:"title"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
