package stores

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zettadam/adamz-api-go/internal/types"
)

type EventStore struct {
	DB *pgxpool.Pool
}

func (s *EventStore) ReadLatest(limit int) ([]types.Event, error) {
	rows, err := s.DB.Query(
		context.Background(),
		`SELECT * FROM events
      ORDER BY created_at DESC
      LIMIT $1`,
		limit)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByName[types.Event])
}

func (s *EventStore) CreateOne(d types.EventRequest) (types.Event, error) {
	result, err := s.DB.Query(context.Background(),
		`INSERT INTO events (
      title, description, start_time, end_time
    ) VALUES (
      $1, $2, $3, $4
    ) RETURNING *`,
		d.Title, d.Description, d.StartTime, d.EndTime)
	if err != nil {
		return types.Event{}, err
	}
	return pgx.CollectOneRow(result, pgx.RowToStructByPos[types.Event])
}

func (s *EventStore) ReadOne(id int64) (types.Event, error) {
	rows, err := s.DB.Query(
		context.Background(),
		`SELECT * FROM events WHERE id = $1`,
		id)
	if err != nil {
		return types.Event{}, err
	}
	return pgx.CollectOneRow(rows, pgx.RowToStructByPos[types.Event])
}

func (s *EventStore) UpdateOne(
	id int64,
	d types.EventRequest,
) (types.Event, error) {
	result, err := s.DB.Query(
		context.Background(),
		`UPDATE events SET (
      title = $2,
      description = $3,
      start_time = $4,
      end_time = $5,
      updated_at = NOW()
    ) WHERE id = $1`,
		id,
		d.Title,
		d.Description,
		d.StartTime,
		d.EndTime,
	)
	if err != nil {
		return types.Event{}, err
	}
	return pgx.CollectOneRow(result, pgx.RowToStructByPos[types.Event])
}

func (s *EventStore) DeleteOne(id int64) (int64, error) {
	result, err := s.DB.Exec(
		context.Background(),
		`DELETE FROM events WHERE id = $1`,
		id)
	return result.RowsAffected(), err
}
