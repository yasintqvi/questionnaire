package port

import (
	"github.com/google/uuid"
	"questionare/internal/core/application/dto/request"
	"questionare/internal/core/application/dto/response"
)

type QuestionnaireService interface {
	GetAllQuestionnaires() ([]*response.QuestionnaireResponse, error)
	FindQuestionnaireById(id uuid.UUID) (*response.QuestionnaireResponse, error)
	CreateQuestionnaire(request *request.QuestionnaireCreateRequest) (*response.QuestionnaireResponse, error)
	UpdateQuestionnaire(id uuid.UUID, request *request.QuestionnaireUpdateRequest) (*response.QuestionnaireResponse, error)
	DeleteQuestionnaire(id uuid.UUID) error
}
