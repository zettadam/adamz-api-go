package stores

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zettadam/adamz-api-go/internal/types"
)

type TaskStore struct {
	DB *pgxpool.Pool
}

func (s *TaskStore) ReadLatest(limit int64) ([]types.Task, error) {
	result, err := s.DB.Query(
		context.Background(),
		`SELECT * FROM tasks
      ORDER BY created_at DESC
      LIMIT $1`,
		limit)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(result, pgx.RowToStructByName[types.Task])
}

func (s *TaskStore) CreateOne(d types.TaskRequest) (types.Task, error) {
	result, err := s.DB.Query(
		context.Background(),
		`INSERT INTO tasks (
      task_id, title, description
    ) VALUES (
      $1, $2, $3
    ) RETURNING *`,
		d.TaskId, d.Title, d.Description)
	if err != nil {
		return types.Task{}, err
	}
	return pgx.CollectOneRow(result, pgx.RowToStructByPos[types.Task])
}

func (s *TaskStore) ReadOne(id int64) (types.Task, error) {
	result, err := s.DB.Query(
		context.Background(),
		`SELECT * FROM tasks WHERE id = $1`,
		id)
	if err != nil {
		return types.Task{}, err
	}
	return pgx.CollectOneRow(result, pgx.RowToStructByPos[types.Task])
}

func (s *TaskStore) UpdateOne(
	id int64,
	d types.TaskRequest,
) (types.Task, error) {
	result, err := s.DB.Query(
		context.Background(),
		`UPDATE tasks SET (
      task_id = $2
      title = $3,
      description = $4,
      updated_at = NOW()
    ) WHERE id = $1
    RETURNING *`,
		id, d.TaskId, d.Title, d.Description)
	if err != nil {
		return types.Task{}, err
	}
	return pgx.CollectOneRow(result, pgx.RowToStructByPos[types.Task])
}

func (s *TaskStore) DeleteOne(id int64) (int64, error) {
	result, err := s.DB.Exec(
		context.Background(),
		`DELETE FROM tasks WHERE id = $1`,
		id)
	return result.RowsAffected(), err
}
