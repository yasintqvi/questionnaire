package http

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	requestDTO "questionare/internal/core/application/dto/request"
	"questionare/internal/core/application/inbound/adapter/service"
)

type QuestionnaireHandler struct {
	service *service.QuestionnaireService
}

func (questionnaireHandler QuestionnaireHandler) GetAllQuestionnaires(writer http.ResponseWriter, request *http.Request) {

	questionnaires, err := questionnaireHandler.service.GetAllQuestionnaires()

	if err != nil {
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusOK)

	err = json.NewEncoder(writer).Encode(questionnaires)

	if err != nil {
		panic(err)
	}
}

func (questionnaireHandler QuestionnaireHandler) FindQuestionnaireById(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	idStr := vars["id"]

	id, err := uuid.Parse(idStr)

	if err != nil {
		panic(err)
	}

	questionnaires, err := questionnaireHandler.service.FindQuestionnaireById(id)

	if err != nil {
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusOK)

	err = json.NewEncoder(writer).Encode(questionnaires)

	if err != nil {
		panic(err)
	}
}

func (questionnaireHandler QuestionnaireHandler) CreateQuestionnaire(writer http.ResponseWriter, request *http.Request) {

	var questionnaireRequest requestDTO.QuestionnaireCreateRequest

	err := json.NewDecoder(request.Body).Decode(&questionnaireRequest)

	if err != nil {
		panic(err)
	}

	questionnaire, err := questionnaireHandler.service.CreateQuestionnaire(questionnaireRequest)

	if err != nil {
		panic(err)
	}

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(writer).Encode(questionnaire)

	if err != nil {
		panic(err)
	}
}

func NewQuestionnaireHandler(service *service.QuestionnaireService) *QuestionnaireHandler {
	return &QuestionnaireHandler{service: service}
}
