package http

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	requestDTO "questionare/internal/core/application/dto/request"
	"questionare/internal/core/application/inbound/adapter/service"
)

type QuestionHandler struct {
	service *service.QuestionService
}

func (questionHandler QuestionHandler) GetAllQuestions(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	idStr := vars["questionnaire_id"]

	id, err := uuid.Parse(idStr)

	if err != nil {
		panic(err)
	}

	questions, err := questionHandler.service.GetAllQuestions(id)

	if err != nil {
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusOK)

	err = json.NewEncoder(writer).Encode(questions)

	if err != nil {
		panic(err)
	}
}

func (questionHandler QuestionHandler) FindQuestionById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	idStr := vars["question_id"]

	id, err := uuid.Parse(idStr)

	if err != nil {
		panic(err)
	}

	question, err := questionHandler.service.FindQuestionById(id)

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusOK)

	err = json.NewEncoder(writer).Encode(question)

	if err != nil {
		panic(err)
	}
}

func (questionHandler QuestionHandler) CreateQuestion(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	idStr := vars["questionnaire_id"]

	id, err := uuid.Parse(idStr)

	if err != nil {
		panic(err)
	}

	var questionRequest requestDTO.QuestionCreateRequest

	err = json.NewDecoder(request.Body).Decode(&questionRequest)

	if err != nil {
		panic(err)
	}

	question, err := questionHandler.service.CreateQuestion(id, &questionRequest)

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusOK)

	err = json.NewEncoder(writer).Encode(question)

	if err != nil {
		panic(err)
	}
}

func (questionHandler QuestionHandler) UpdateQuestion(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	idStr := vars["question_id"]

	id, err := uuid.Parse(idStr)

	if err != nil {
		panic(err)
	}

	var questionRequest requestDTO.QuestionUpdateRequest

	err = json.NewDecoder(request.Body).Decode(&questionRequest)

	if err != nil {
		panic(err)
	}

	question, err := questionHandler.service.UpdateQuestion(id, &questionRequest)

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusOK)

	err = json.NewEncoder(writer).Encode(question)

	if err != nil {
		panic(err)
	}
}

func (questionHandler QuestionHandler) DeleteQuestion(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	idStr := vars["question_id"]

	id, err := uuid.Parse(idStr)

	if err != nil {
		panic(err)
	}

	err = questionHandler.service.DeleteQuestion(id)

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusNoContent)
}

func NewQuestionHandler(service *service.QuestionService) *QuestionHandler {
	return &QuestionHandler{service: service}
}
