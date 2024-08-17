package stores

import (
	"database/sql"
	"errors"
	"time"

	"github.com/zettadam/adamz-api-go/internal/models"
)

type NoteStore struct {
	DB *sql.DB
}

func (s *NoteStore) readLatest(limit int) ([]*models.Note, error) {
	rows, err := s.DB.Query(
		`SELECT * FROM notes 
      ORDER BY published_at DESC, created_at DESC 
      LIMIT $1`,
		limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	data := []*models.Note{}
	for rows.Next() {
		d := &models.Note{}
		err = rows.Scan(
			&d.Id,
			&d.Title,
			&d.Body,
			&d.Significance,
			&d.PublishedAt,
			&d.Tags,
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

func (s *NoteStore) createOne(
	title string,
	body string,
	significance int,
	publishedAt int,
	tags []string,
) (int64, error) {
	result, err := s.DB.Exec(
		`INSERT INTO notes (
      title, body, significance, published_at, tags
    ) VALUES (
      $1, $2, $3, $4, $5
    )`,
		title, body, significance, publishedAt, tags)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *NoteStore) readOne(id int64) (*models.Note, error) {
	d := &models.Note{}

	err := s.DB.QueryRow(
		`SELECT * FROM notes WHERE id = $1`, id).Scan(
		&d.Id,
		&d.Title,
		&d.Body,
		&d.Significance,
		&d.PublishedAt,
		&d.Tags,
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

func (s *NoteStore) updateOne(
	id int64,
	title string,
	body string,
	significance int,
	publishedAt time.Time,
	tags []string,
) (int64, error) {
	result, err := s.DB.Exec(
		`UPDATE notes SET (
      title = $2,
      body = $3,
      significance = $4,
      published_at = $5,
      tags = $6,
      updated_at = NOW()
    ) WHERE id = $1`,
		id, title, body, significance, publishedAt, tags)
	if err != nil {
		return 0, err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return n, nil
}

func (s *NoteStore) deleteOne(id int64) (int64, error) {
	result, err := s.DB.Exec(`DELETE FROM notes WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return n, nil
}
