package adapter

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"questionare/internal/core/domain"
	"time"
)

type QuestionnaireRepositoryMySql struct {
	db *sql.DB
}

func (questionnaireRepo QuestionnaireRepositoryMySql) GetAll() ([]*domain.Questionnaire, error) {

	query := "SELECT id, title, description, status, start_time, end_time, created_at, updated_at, deleted_at" +
		" FROM questionnaires WHERE deleted_at IS NULL ORDER BY created_at DESC"

	rows, err := questionnaireRepo.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	fmt.Println(rows)
	var questionnaires []*domain.Questionnaire
	for rows.Next() {
		var questionnaire domain.Questionnaire
		err = rows.Scan(&questionnaire.ID, &questionnaire.Title, &questionnaire.Description, &questionnaire.Status,
			&questionnaire.StartTime, &questionnaire.EndTime, &questionnaire.CreatedAt, &questionnaire.UpdatedAt,
			&questionnaire.DeletedAt)

		if err != nil {
			return nil, err
		}

		questionnaires = append(questionnaires, &questionnaire)
	}

	return questionnaires, nil
}

func (questionnaireRepo QuestionnaireRepositoryMySql) Save(questionnaire *domain.Questionnaire) (*domain.Questionnaire, error) {
	query := "INSERT INTO questionnaires (id, title, description, status, start_time, end_time, created_at, updated_at)" +
		"  VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	questionnaire.ID = uuid.New()
	questionnaire.CreatedAt = time.Now()
	questionnaire.UpdatedAt = time.Now()

	_, err := questionnaireRepo.db.Exec(query, &questionnaire.ID, &questionnaire.Title,
		&questionnaire.Description, &questionnaire.Status, questionnaire.StartTime, questionnaire.EndTime,
		questionnaire.CreatedAt, questionnaire.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return questionnaire, nil
}

func (questionnaireRepo QuestionnaireRepositoryMySql) FindById(id uuid.UUID) (*domain.Questionnaire, error) {

	query := "SELECT id, title, description, status, start_time, end_time, created_at, updated_at, deleted_at" +
		" FROM questionnaires WHERE id = ? AND deleted_at IS NULL"

	row := questionnaireRepo.db.QueryRow(query, id)

	var questionnaire domain.Questionnaire

	err := row.Scan(&questionnaire.ID, &questionnaire.Title, &questionnaire.Description, &questionnaire.Status,
		&questionnaire.StartTime, &questionnaire.EndTime, &questionnaire.CreatedAt, &questionnaire.UpdatedAt,
		&questionnaire.DeletedAt)

	if err != nil {
		return nil, err
	}

	return &questionnaire, nil
}

func (questionnaireRepo QuestionnaireRepositoryMySql) Update(questionnaire *domain.Questionnaire) (*domain.Questionnaire, error) {
	query := "UPDATE questionnaires SET title =?, description=?, status=?, start_time=?, end_time=? WHERE id = ?" +
		" AND deleted_at IS NULL"

	_, err := questionnaireRepo.db.Exec(query, questionnaire.Title, questionnaire.Description,
		questionnaire.Status, questionnaire.StartTime, questionnaire.EndTime, questionnaire.ID)

	if err != nil {
		return nil, err
	}

	return questionnaire, nil
}

func (questionnaireRepo QuestionnaireRepositoryMySql) DeleteById(id uuid.UUID) error {
	query := "UPDATE questionnaires SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL"

	_, err := questionnaireRepo.db.Exec(query, time.Now(), id)

	if err != nil {
		return err
	}

	return nil
}

func NewMysqlQuestionnaireRepository(db *sql.DB) QuestionnaireRepositoryMySql {
	return QuestionnaireRepositoryMySql{db: db}
}
