package port

import "github.com/google/uuid"

type QuestionService interface {
	GetAllQuestions(questionnaireId uuid.UUID)
	FindQuestionById(id uuid.UUID)
	CreateQuestion(questionnaireId uuid.UUID)
	UpdateQuestion(uuid uuid.UUID)
	DeleteQuestion(uuid uuid.UUID)
}
