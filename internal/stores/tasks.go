package stores

import (
	"database/sql"
	"errors"

	"github.com/zettadam/adamz-api-go/internal/models"
)

type TaskStore struct {
	DB *sql.DB
}

func (s *TaskStore) readLatest(limit int64) ([]*models.Task, error) {
	rows, err := s.DB.Query(
		`SELECT * FROM tasks 
      ORDER BY created_at DESC 
      LIMIT $1`,
		limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	data := []*models.Task{}
	for rows.Next() {
		d := &models.Task{}
		err = rows.Scan(
			&d.Id,
			&d.TaskId,
			&d.Title,
			&d.Description,
			&d.CreatedAt,
			&d.UpdatedAt)
		if err != nil {
			return nil, err
		}
		data = append(data, d)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

func (s *TaskStore) createOne(
	taskId int64,
	title string,
	description string,
) (int64, error) {
	result, err := s.DB.Exec(
		`INSERT INTO tasks (
      task_id, title, description
    ) VALUES (
      $1, $2, $3, $4, $5, $6
    )`,
		taskId, title, description)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *TaskStore) readOne(id int64) (*models.Task, error) {
	d := &models.Task{}

	err := s.DB.QueryRow(
		`SELECT * FROM tasks WHERE id = $1`, id).Scan(
		&d.Id,
		&d.TaskId,
		&d.Title,
		&d.Description,
		&d.CreatedAt,
		&d.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return d, nil
}

func (s *TaskStore) updateOne(
	id int64,
	taskId int64,
	title string,
	description string,
) (int64, error) {
	result, err := s.DB.Exec(
		`UPDATE tasks SET (
      task_id = $2
      title = $3,
      description = $4,
      updated_at = NOW()
    ) WHERE id = $1`,
		id, taskId, title, description)
	if err != nil {
		return 0, err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return n, nil
}

func (s *TaskStore) deleteOne(id int64) (int64, error) {
	result, err := s.DB.Exec(`DELETE FROM tasks WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return n, nil
}
