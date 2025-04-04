package port

import (
	"github.com/google/uuid"
	"questionare/internal/core/application/dto/request"
	"questionare/internal/core/application/dto/response"
)

type QuestionService interface {
	GetAllQuestions(questionnaireId uuid.UUID) ([]*response.QuestionResponse, error)
	FindQuestionById(questionId uuid.UUID) (*response.QuestionResponse, error)
	CreateQuestion(questionnaireId uuid.UUID, request *request.QuestionCreateRequest) (*response.QuestionResponse, error)
	UpdateQuestion(questionId uuid.UUID, request *request.QuestionUpdateRequest) (*response.QuestionResponse, error)
	DeleteQuestion(questionId uuid.UUID) error
}
