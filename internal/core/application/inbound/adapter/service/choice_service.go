package service

import (
	"github.com/google/uuid"
	"questionare/internal/core/application/dto/request"
	"questionare/internal/core/application/dto/response"
	"questionare/internal/core/application/outbound/port"
	"questionare/internal/core/domain"
)

type ChoiceService struct {
	choiceRepository port.ChoiceRepository
}

func (choiceService ChoiceService) GetAllChoices(questionId uuid.UUID) ([]*response.ChoiceResponse, error) {
	choices, err := choiceService.choiceRepository.GetAll(questionId)

	if err != nil {
		return nil, err
	}

	responseChoices := make([]*response.ChoiceResponse, len(choices))

	for i, choice := range choices {
		responseChoices[i] = &response.ChoiceResponse{
			ID:         choice.ID,
			QuestionId: questionId,
			Value:      choice.Value,
			CreatedAt:  choice.CreatedAt,
			UpdatedAt:  choice.UpdatedAt,
		}
	}

	return responseChoices, nil
}

func (choiceService ChoiceService) FindChoiceById(choiceId uuid.UUID) (*response.ChoiceResponse, error) {

	choice, err := choiceService.choiceRepository.FindById(choiceId)

	if err != nil {
		return nil, err
	}

	responseChoice := &response.ChoiceResponse{
		ID:         choice.ID,
		QuestionId: choice.QuestionId,
		Value:      choice.Value,
		CreatedAt:  choice.CreatedAt,
		UpdatedAt:  choice.UpdatedAt,
	}

	return responseChoice, nil
}

func (choiceService ChoiceService) CreateChoice(questionId uuid.UUID, request *request.ChoiceCreateRequest) (*response.ChoiceResponse, error) {

	question := domain.Choice{
		QuestionId: questionId,
		Value:      request.Value,
	}

	newChoice, err := choiceService.choiceRepository.Save(&question)

	if err != nil {
		return nil, err
	}

	responseChoice := &response.ChoiceResponse{
		ID:         newChoice.ID,
		QuestionId: newChoice.QuestionId,
		Value:      newChoice.Value,
		CreatedAt:  newChoice.CreatedAt,
		UpdatedAt:  newChoice.UpdatedAt,
	}

	return responseChoice, nil
}

func (choiceService ChoiceService) UpdateChoice(questionId uuid.UUID, request *request.ChoiceUpdateRequest) (*response.ChoiceResponse, error) {
	question, err := choiceService.choiceRepository.FindById(questionId)

	question.Value = request.Value

	if err != nil {
		return nil, err
	}

	updatedChoice, err := choiceService.choiceRepository.Update(question)

	if err != nil {
		return nil, err
	}

	responseChoice := &response.ChoiceResponse{
		ID:         updatedChoice.ID,
		QuestionId: question.QuestionId,
		Value:      updatedChoice.Value,
		CreatedAt:  updatedChoice.CreatedAt,
		UpdatedAt:  updatedChoice.UpdatedAt,
	}

	return responseChoice, nil
}

func (choiceService ChoiceService) DeleteChoice(questionId uuid.UUID) error {
	_, err := choiceService.choiceRepository.FindById(questionId)

	if err != nil {
		return err
	}

	err = choiceService.choiceRepository.Delete(questionId)

	if err != nil {
		return err
	}

	return nil
}

func NewChoiceService(choiceRepository port.ChoiceRepository) *ChoiceService {
	return &ChoiceService{choiceRepository: choiceRepository}
}
