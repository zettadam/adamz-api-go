package stores

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zettadam/adamz-api-go/internal/models"
)

type EventStore struct {
	DB *pgxpool.Pool
}

func (s *EventStore) ReadLatest(limit int) ([]models.Event, error) {
	rows, err := s.DB.Query(
		context.Background(),
		`SELECT * FROM events
      ORDER BY created_at DESC
      LIMIT $1`,
		limit)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Event])
}

func (s *EventStore) CreateOne(
	title string,
	description string,
	startTime time.Time,
	endTime time.Time,
) (int64, error) {
	var id int64 = 0
	err := s.DB.QueryRow(context.Background(),
		`INSERT INTO events (
      title, description, start_time, end_time
    ) VALUES (
      $1, $2, $3, $4
    ) RETURNING id`,
		title, description, startTime, endTime).Scan(&id)
	return id, err
}

func (s *EventStore) ReadOne(id int64) (models.Event, error) {
	rows, _ := s.DB.Query(
		context.Background(),
		`SELECT * FROM events WHERE id = $1`,
		id)
	return pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.Event])
}

func (s *EventStore) UpdateOne(
	id int64,
	title string,
	description string,
	startTime time.Time,
	endTime time.Time,
) (int64, error) {
	result, err := s.DB.Exec(context.Background(),
		`UPDATE events SET (
      title = $2,
      description = $3,
      start_time = $4,
      end_time = $5,
      updated_at = NOW()
    ) WHERE id = $1`,
		id, title, description, startTime, endTime)
	return result.RowsAffected(), err
}

func (s *EventStore) DeleteOne(id int64) (int64, error) {
	result, err := s.DB.Exec(
		context.Background(),
		`DELETE FROM events WHERE id = $1`,
		id)
	return result.RowsAffected(), err
}
