package port

import (
	"github.com/google/uuid"
	"questionare/internal/core/domain"
)

type QuestionnaireRepository interface {
	GetAll() ([]*domain.Questionnaire, error)
	Create(questionnaire *domain.Questionnaire) (*domain.Questionnaire, error)
	FindById(id uuid.UUID) (*domain.Questionnaire, error)
	Update(questionnaire *domain.Questionnaire) (*domain.Questionnaire, error)
	DeleteById(id uuid.UUID) error
}
