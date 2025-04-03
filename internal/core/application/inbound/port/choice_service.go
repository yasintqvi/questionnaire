package port

import "github.com/google/uuid"

type ChoiceService interface {
	GetAllChoices()
	GetChoiceById(id uuid.UUID)
	CreateChoice()
	UpdateChoice(id uuid.UUID)
	DeleteChoice(uuid uuid.UUID)
}
