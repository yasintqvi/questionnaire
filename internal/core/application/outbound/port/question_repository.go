package port

import (
	"github.com/google/uuid"
	"questionare/internal/core/domain"
)

type QuestionRepository interface {
	GetAll(questionnaireId uuid.UUID) ([]*domain.Question, error)
	FindById(uuid uuid.UUID) (*domain.Question, error)
	Save(question *domain.Question) (*domain.Question, error)
	Update(question *domain.Question) (*domain.Question, error)
	Delete(uuid uuid.UUID) error
}
