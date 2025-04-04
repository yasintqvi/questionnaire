package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	httpHandler "questionare/internal/core/application/inbound/adapter/http"
	"questionare/internal/core/application/inbound/adapter/service"
	"questionare/internal/core/application/outbound/adapter"
	"time"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
	}

	for _, prop := range envProps {
		if os.Getenv(prop) == "" {
			log.Fatalf("Environment variable %s is not set", prop)
		}
	}
}

func main() {
	fmt.Println("Application Started ...")
	sanityCheck()

	router := mux.NewRouter()

	db := getClientDb()

	questionnaireRepo := adapter.NewMysqlQuestionnaireRepository(db)
	questionRepo := adapter.NewMysqlQuestionRepository(db)
	choiceRepo := adapter.NewMysqlChoiceRepository(db)

	questionnaireHandler := httpHandler.NewQuestionnaireHandler(service.NewQuestionnaireService(questionnaireRepo))
	questionHandler := httpHandler.NewQuestionHandler(service.NewQuestionService(questionRepo))
	choiceHandler := httpHandler.NewChoiceHandler(service.NewChoiceService(choiceRepo))

	router.HandleFunc("/api/questionnaires", questionnaireHandler.GetAllQuestionnaires).Methods(http.MethodGet)
	router.HandleFunc("/api/questionnaires/{id}", questionnaireHandler.FindQuestionnaireById).Methods(http.MethodGet)
	router.HandleFunc("/api/questionnaires/create", questionnaireHandler.CreateQuestionnaire).Methods(http.MethodPost)
	router.HandleFunc("/api/questionnaires/{id}", questionnaireHandler.UpdateQuestionnaire).Methods(http.MethodPut)
	router.HandleFunc("/api/questionnaires/{id}", questionnaireHandler.DeleteQuestionnaire).Methods(http.MethodDelete)

	router.HandleFunc("/api/questionnaires/{questionnaire_id}/questions", questionHandler.GetAllQuestions).Methods(http.MethodGet)
	router.HandleFunc("/api/questionnaires/{questionnaire_id}/questions/{question_id}", questionHandler.FindQuestionById).Methods(http.MethodGet)
	router.HandleFunc("/api/questionnaires/{questionnaire_id}/questions", questionHandler.CreateQuestion).Methods(http.MethodPost)
	router.HandleFunc("/api/questionnaires/{questionnaire_id}/questions/{question_id}", questionHandler.UpdateQuestion).Methods(http.MethodPut)
	router.HandleFunc("/api/questionnaires/{questionnaire_id}/questions/{question_id}", questionHandler.DeleteQuestion).Methods(http.MethodDelete)

	router.HandleFunc("/api/questions/{question_id}/choices", choiceHandler.GetAllChoices).Methods(http.MethodGet)
	router.HandleFunc("/api/questions/{question_id}/choices/{choice_id}", choiceHandler.FindChoiceById).Methods(http.MethodGet)
	router.HandleFunc("/api/questions/{question_id}/choices", choiceHandler.CreateChoice).Methods(http.MethodPost)
	router.HandleFunc("/api/questions/{question_id}/choices/{choice_id}", choiceHandler.UpdateChoice).Methods(http.MethodPut)
	router.HandleFunc("/api/questions/{question_id}/choices/{choice_id}", choiceHandler.DeleteChoice).Methods(http.MethodDelete)

	err := http.ListenAndServe(os.Getenv("SERVER_URL")+":"+os.Getenv("SERVER_PORT"), router)

	if err != nil {
		panic(err)
	}
}

func getClientDb() *sql.DB {
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	client, err := sql.Open("mysql", dbInfo)
	if err != nil {
		panic(fmt.Sprintf("An error occurred while conneting to the database: %v", err))
	}

	err = client.Ping()
	if err != nil {
		panic(fmt.Sprintf("Error pinging database: %v", err))
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
