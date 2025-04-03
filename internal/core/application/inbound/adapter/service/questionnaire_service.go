package service

import (
	"github.com/google/uuid"
	"log"
	"questionare/internal/core/application/dto/request"
	"questionare/internal/core/application/dto/response"
	"questionare/internal/core/application/outbound/port"
	"questionare/internal/core/domain"
	"time"
)

type QuestionnaireService struct {
	questionnaireRepository port.QuestionnaireRepository
}

func (questionnaireService QuestionnaireService) GetAllQuestionnaires() ([]*response.QuestionnaireResponse, error) {
	questionnaires, err := questionnaireService.questionnaireRepository.GetAll()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	responseQuestionnaires := make([]*response.QuestionnaireResponse, 0, len(questionnaires))

	for _, questionnaire := range questionnaires {
		responseQuestionnaires = append(responseQuestionnaires, &response.QuestionnaireResponse{
			ID:          questionnaire.ID,
			Title:       questionnaire.Title,
			Description: questionnaire.Description,
			StartTime:   questionnaire.StartTime.Format(time.DateTime),
			EndTime:     questionnaire.EndTime.Format(time.DateTime),
			Status:      questionnaire.Status,
			CreatedAt:   questionnaire.CreatedAt.Format(time.DateTime),
			UpdatedAt:   questionnaire.UpdatedAt.Format(time.DateTime),
		})
	}

	return responseQuestionnaires, nil
}

func (questionnaireService QuestionnaireService) FindQuestionnaireById(id uuid.UUID) (*response.QuestionnaireResponse, error) {

	questionnaire, err := questionnaireService.questionnaireRepository.FindById(id)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &response.QuestionnaireResponse{
		ID:          questionnaire.ID,
		Title:       questionnaire.Title,
		Description: questionnaire.Description,
		StartTime:   questionnaire.StartTime.Format(time.DateTime),
		EndTime:     questionnaire.EndTime.Format(time.DateTime),
		Status:      questionnaire.Status,
		CreatedAt:   questionnaire.CreatedAt.Format(time.DateTime),
		UpdatedAt:   questionnaire.UpdatedAt.Format(time.DateTime),
	}, nil
}

func (questionnaireService QuestionnaireService) CreateQuestionnaire(request *request.QuestionnaireCreateRequest) (*response.QuestionnaireResponse, error) {

	startTime, err := time.Parse(time.DateTime, request.StartTime)

	endTime, err := time.Parse(time.DateTime, request.EndTime)

	if err != nil {
		return nil, err
	}

	questionnaire := domain.Questionnaire{
		Title:       request.Title,
		Description: request.Description,
		StartTime:   startTime,
		EndTime:     endTime,
	}

	newQuestionnaire, err := questionnaireService.questionnaireRepository.Save(&questionnaire)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	questionnaireResponse := &response.QuestionnaireResponse{
		ID:          newQuestionnaire.ID,
		Title:       newQuestionnaire.Title,
		Description: newQuestionnaire.Description,
		StartTime:   newQuestionnaire.StartTime.Format(time.DateTime),
		EndTime:     newQuestionnaire.EndTime.Format(time.DateTime),
		Status:      newQuestionnaire.Status,
		CreatedAt:   newQuestionnaire.CreatedAt.Format(time.DateTime),
		UpdatedAt:   newQuestionnaire.UpdatedAt.Format(time.DateTime),
	}

	return questionnaireResponse, nil
}

func (questionnaireService QuestionnaireService) UpdateQuestionnaire(id uuid.UUID, request *request.QuestionnaireUpdateRequest) (*response.QuestionnaireResponse, error) {
	startTime, err := time.Parse(time.DateTime, request.StartTime)
	endTime, err := time.Parse(time.DateTime, request.EndTime)

	if err != nil {
		return nil, err
	}

	questionnaire := domain.Questionnaire{
		ID:          id,
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
		StartTime:   startTime,
		EndTime:     endTime,
	}

	updatedQuestionnaire, err := questionnaireService.questionnaireRepository.Update(&questionnaire)

	if err != nil {
		log.Fatal(err)
	}

	questionnaireResponse := &response.QuestionnaireResponse{
		ID:          updatedQuestionnaire.ID,
		Title:       updatedQuestionnaire.Title,
		Description: updatedQuestionnaire.Description,
		StartTime:   updatedQuestionnaire.StartTime.Format(time.DateTime),
		EndTime:     updatedQuestionnaire.EndTime.Format(time.DateTime),
		Status:      updatedQuestionnaire.Status,
		CreatedAt:   updatedQuestionnaire.CreatedAt.Format(time.DateTime),
		UpdatedAt:   updatedQuestionnaire.UpdatedAt.Format(time.DateTime),
	}

	return questionnaireResponse, nil
}

func (questionnaireService QuestionnaireService) DeleteQuestionnaire(id uuid.UUID) error {

	questionnaire, err := questionnaireService.questionnaireRepository.FindById(id)

	if err != nil {
		return err
	}

	if err := questionnaireService.questionnaireRepository.DeleteById(questionnaire.ID); err != nil {
		log.Fatal(err)
	}

	return nil
}

func NewQuestionnaireService(repository port.QuestionnaireRepository) *QuestionnaireService {
	return &QuestionnaireService{repository}
}
