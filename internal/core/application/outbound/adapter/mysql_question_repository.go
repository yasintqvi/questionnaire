package adapter

import (
	"database/sql"
	"github.com/google/uuid"
	"questionare/internal/core/domain"
	"time"
)

type MysqlQuestionRepository struct {
	db *sql.DB
}

func (questionRepo MysqlQuestionRepository) GetAll(questionnaireId uuid.UUID) ([]*domain.Question, error) {

	query := `
		SELECT id, questionnaire_id, title, created_at, updated_at, deleted_at FROM questions q
		                                                                       WHERE q.questionnaire_id = ?
		                                                                         AND q.deleted_at IS NULL`

	rows, err := questionRepo.db.Query(query, questionnaireId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var questions []*domain.Question

	for rows.Next() {
		var question domain.Question

		err := rows.Scan(&question.ID, &question.QuestionnaireId, &question.Title, &question.CreatedAt,
			&question.UpdatedAt, &question.DeletedAt)

		if err != nil {
			return nil, err
		}

		questions = append(questions, &question)
	}

	return questions, nil
}

func (questionRepo MysqlQuestionRepository) FindById(id uuid.UUID) (*domain.Question, error) {
	query := "SELECT id, questionnaire_id, title, created_at, updated_at, deleted_at FROM questions WHERE id = ? AND deleted_at IS NULL"

	row := questionRepo.db.QueryRow(query, id)

	var question domain.Question

	err := row.Scan(&question.ID, &question.Title, &question.CreatedAt, &question.UpdatedAt, &question.DeletedAt)

	if err != nil {
		return nil, err
	}

	return &question, nil
}

func (questionRepo MysqlQuestionRepository) Create(question *domain.Question) (*domain.Question, error) {

	query := "INSERT INTO questions (id, questionnaire_id, title, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"

	question.ID = uuid.New()

	_, err := questionRepo.db.Exec(query, question.ID, question.QuestionnaireId, question.Title, time.Now(), time.Now())

	if err != nil {
		return nil, err
	}

	return question, nil
}

func (questionRepo MysqlQuestionRepository) Update(question *domain.Question) (*domain.Question, error) {
	query := "UPDATE questions SET title=?, questionnaire_id=? WHERE id=?"

	_, err := questionRepo.db.Exec(query, &question.Title, &question.QuestionnaireId)

	if err != nil {
		return nil, err
	}

	return question, nil
}

func (questionRepo MysqlQuestionRepository) Delete(id uuid.UUID) error {
	query := "UPDATE questions SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL"

	_, err := questionRepo.db.Exec(query, time.Now(), id)

	if err != nil {
		return err
	}

	return nil
}

func NewMysqlQuestionRepository(db *sql.DB) MysqlQuestionRepository {
	return MysqlQuestionRepository{db: db}
}
