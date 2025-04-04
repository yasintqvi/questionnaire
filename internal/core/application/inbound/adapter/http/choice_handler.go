package http

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	requestDTO "questionare/internal/core/application/dto/request"
	"questionare/internal/core/application/inbound/adapter/service"
)

type ChoiceHandler struct {
	service *service.ChoiceService
}

func (choiceHandler ChoiceHandler) GetAllChoices(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	idStr := vars["question_id"]

	id, err := uuid.Parse(idStr)

	if err != nil {
		panic(err)
	}

	questions, err := choiceHandler.service.GetAllChoices(id)

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

func (choiceHandler ChoiceHandler) FindChoiceById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	idStr := vars["choice_id"]

	id, err := uuid.Parse(idStr)

	if err != nil {
		panic(err)
	}

	question, err := choiceHandler.service.FindChoiceById(id)

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusOK)

	err = json.NewEncoder(writer).Encode(question)

	if err != nil {
		panic(err)
	}
}

func (choiceHandler ChoiceHandler) CreateChoice(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	idStr := vars["question_id"]

	id, err := uuid.Parse(idStr)

	if err != nil {
		panic(err)
	}

	var questionRequest requestDTO.ChoiceCreateRequest

	err = json.NewDecoder(request.Body).Decode(&questionRequest)

	if err != nil {
		panic(err)
	}

	question, err := choiceHandler.service.CreateChoice(id, &questionRequest)

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusOK)

	err = json.NewEncoder(writer).Encode(question)

	if err != nil {
		panic(err)
	}
}

func (choiceHandler ChoiceHandler) UpdateChoice(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	idStr := vars["choice_id"]

	id, err := uuid.Parse(idStr)

	if err != nil {
		panic(err)
	}

	var questionRequest requestDTO.ChoiceUpdateRequest

	err = json.NewDecoder(request.Body).Decode(&questionRequest)

	if err != nil {
		panic(err)
	}

	question, err := choiceHandler.service.UpdateChoice(id, &questionRequest)

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusOK)

	err = json.NewEncoder(writer).Encode(question)

	if err != nil {
		panic(err)
	}
}

func (choiceHandler ChoiceHandler) DeleteChoice(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	idStr := vars["choice_id"]

	id, err := uuid.Parse(idStr)

	if err != nil {
		panic(err)
	}

	err = choiceHandler.service.DeleteChoice(id)

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusNoContent)
}

func NewChoiceHandler(service *service.ChoiceService) *ChoiceHandler {
	return &ChoiceHandler{service: service}
}
