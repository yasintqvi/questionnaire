package adapter

import (
	"database/sql"
	"github.com/google/uuid"
	"questionare/internal/core/domain"
	"time"
)

type ChoiceRepositoryMySql struct {
	db *sql.DB
}

func (choiceRepo ChoiceRepositoryMySql) GetAll(questionId uuid.UUID) ([]*domain.Choice, error) {

	query := `
		SELECT id, choice_id, value, created_at, updated_at, deleted_at FROM choices q
		                                                                       WHERE choice_id = ?
		                                                                         AND deleted_at IS NULL`

	rows, err := choiceRepo.db.Query(query, questionId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var choices []*domain.Choice

	for rows.Next() {
		var choice domain.Choice

		err := rows.Scan(&choice.ID, &choice.QuestionId, &choice.Value, &choice.CreatedAt,
			&choice.UpdatedAt, &choice.DeletedAt)

		if err != nil {
			return nil, err
		}

		choices = append(choices, &choice)
	}

	return choices, nil
}

func (choiceRepo ChoiceRepositoryMySql) FindById(id uuid.UUID) (*domain.Choice, error) {
	query := "SELECT id, choice_id, value, created_at, updated_at, deleted_at FROM choices WHERE id = ? AND deleted_at IS NULL"

	row := choiceRepo.db.QueryRow(query, id)

	var choice domain.Choice

	err := row.Scan(&choice.ID, &choice.Value, &choice.CreatedAt, &choice.UpdatedAt, &choice.DeletedAt)

	if err != nil {
		return nil, err
	}

	return &choice, nil
}

func (choiceRepo ChoiceRepositoryMySql) Save(choice *domain.Choice) (*domain.Choice, error) {

	query := "INSERT INTO choices (id, choice_id, value, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"

	choice.ID = uuid.New()

	_, err := choiceRepo.db.Exec(query, choice.ID, choice.QuestionId, choice.Value, time.Now(), time.Now())

	if err != nil {
		return nil, err
	}

	return choice, nil
}

func (choiceRepo ChoiceRepositoryMySql) Update(choice *domain.Choice) (*domain.Choice, error) {
	query := "UPDATE choices SET title=?, choicenaire_id=? WHERE id=?"

	_, err := choiceRepo.db.Exec(query, &choice.Value, &choice.QuestionId)

	if err != nil {
		return nil, err
	}

	return choice, nil
}

func (choiceRepo ChoiceRepositoryMySql) Delete(id uuid.UUID) error {
	query := "UPDATE choices SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL"

	_, err := choiceRepo.db.Exec(query, time.Now(), id)

	if err != nil {
		return err
	}

	return nil
}
