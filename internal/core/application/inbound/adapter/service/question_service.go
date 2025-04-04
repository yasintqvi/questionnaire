package service

import (
	"github.com/google/uuid"
	"questionare/internal/core/application/dto/request"
	"questionare/internal/core/application/dto/response"
	"questionare/internal/core/application/outbound/port"
	"questionare/internal/core/domain"
)

type QuestionService struct {
	questionRepository port.QuestionRepository
}

func (questionService QuestionService) GetAllQuestions(questionnaireId uuid.UUID) ([]*response.QuestionResponse, error) {
	questions, err := questionService.questionRepository.GetAll(questionnaireId)

	if err != nil {
		return nil, err
	}

	responseQuestions := make([]*response.QuestionResponse, len(questions))

	for i, question := range questions {
		responseQuestions[i] = &response.QuestionResponse{
			ID:              question.ID,
			QuestionnaireId: questionnaireId,
			Title:           question.Title,
			CreatedAt:       question.CreatedAt,
			UpdatedAt:       question.UpdatedAt,
		}
	}

	return responseQuestions, nil
}

func (questionService QuestionService) FindQuestionById(questionId uuid.UUID) (*response.QuestionResponse, error) {

	question, err := questionService.questionRepository.FindById(questionId)

	if err != nil {
		return nil, err
	}

	responseQuestion := &response.QuestionResponse{
		ID:              question.ID,
		QuestionnaireId: question.QuestionnaireId,
		Title:           question.Title,
		CreatedAt:       question.CreatedAt,
		UpdatedAt:       question.UpdatedAt,
	}

	return responseQuestion, nil
}

func (questionService QuestionService) CreateQuestion(questionnaireId uuid.UUID, request *request.QuestionCreateRequest) (*response.QuestionResponse, error) {

	question := domain.Question{
		QuestionnaireId: questionnaireId,
		Title:           request.Title,
	}

	newQuestion, err := questionService.questionRepository.Save(&question)

	if err != nil {
		return nil, err
	}

	responseQuestion := &response.QuestionResponse{
		ID:              newQuestion.QuestionnaireId,
		QuestionnaireId: newQuestion.QuestionnaireId,
		Title:           newQuestion.Title,
		CreatedAt:       newQuestion.CreatedAt,
		UpdatedAt:       newQuestion.UpdatedAt,
	}

	return responseQuestion, nil
}

func (questionService QuestionService) UpdateQuestion(questionId uuid.UUID, request *request.QuestionUpdateRequest) (*response.QuestionResponse, error) {
	question, err := questionService.questionRepository.FindById(questionId)

	question.Title = request.Title

	if err != nil {
		return nil, err
	}

	updatedQuestion, err := questionService.questionRepository.Update(question)

	if err != nil {
		return nil, err
	}

	responseQuestion := &response.QuestionResponse{
		ID:              updatedQuestion.ID,
		QuestionnaireId: question.QuestionnaireId,
		Title:           updatedQuestion.Title,
		CreatedAt:       updatedQuestion.CreatedAt,
		UpdatedAt:       updatedQuestion.UpdatedAt,
	}

	return responseQuestion, nil
}

func (questionService QuestionService) DeleteQuestion(questionId uuid.UUID) error {
	_, err := questionService.questionRepository.FindById(questionId)

	if err != nil {
		return err
	}

	err = questionService.questionRepository.Delete(questionId)

	if err != nil {
		return err
	}

	return nil
}

func NewQuestionService(questionRepository port.QuestionRepository) *QuestionService {
	return &QuestionService{questionRepository: questionRepository}
}
