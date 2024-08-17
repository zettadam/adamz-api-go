package stores

import (
	"database/sql"
	"errors"
	"time"

	"github.com/zettadam/adamz-api-go/internal/models"
)

type EventStore struct {
	DB *sql.DB
}

func (s *EventStore) readLatest(limit int) ([]*models.Event, error) {
	rows, err := s.DB.Query(
		`SELECT * FROM events 
      ORDER BY created_at DESC 
      LIMIT $1`,
		limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	data := []*models.Event{}
	for rows.Next() {
		d := &models.Event{}
		err = rows.Scan(
			&d.Id,
			&d.Title,
			&d.Description,
			&d.StartTime,
			&d.EndTime,
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

func (s *EventStore) createOne(
	title string,
	description string,
	startTime time.Time,
	endTime time.Time,
) (int64, error) {
	result, err := s.DB.Exec(
		`INSERT INTO events (
      title, description, start_time, end_time
    ) VALUES (
      $1, $2, $3, $4
    )`,
		title, description, startTime, endTime)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *EventStore) readOne(id int64) (*models.Event, error) {
	d := &models.Event{}

	err := s.DB.QueryRow(
		`SELECT * FROM events WHERE id = $1`, id).Scan(
		&d.Id,
		&d.Title,
		&d.Description,
		&d.StartTime,
		&d.EndTime,
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

func (s *EventStore) updateOne(
	id int64,
	title string,
	description string,
	startTime time.Time,
	endTime time.Time,
) (int64, error) {
	result, err := s.DB.Exec(
		`UPDATE events SET (
      title = $2,
      description = $3,
      start_time = $4,
      end_time = $5,
      updated_at = NOW()
    ) WHERE id = $1`,
		id, title, description, startTime, endTime)
	if err != nil {
		return 0, err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return n, nil
}

func (s *EventStore) deleteOne(id int64) (int64, error) {
	result, err := s.DB.Exec(`DELETE FROM events WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return n, nil
}
