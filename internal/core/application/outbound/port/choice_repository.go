package port

import (
	"github.com/google/uuid"
	"questionare/internal/core/domain"
)

type ChoiceRepository interface {
	GetAll(questionId uuid.UUID) ([]*domain.Choice, error)
	FindById(id uuid.UUID) (*domain.Choice, error)
	Create(choice *domain.Choice) (*domain.Choice, error)
	Update(choice *domain.Choice) (*domain.Choice, error)
	Delete(id uuid.UUID) error
}
