package stores

import (
	"database/sql"
	"errors"
	"time"

	"github.com/zettadam/adamz-api-go/internal/models"
)

type LinkStore struct {
	DB *sql.DB
}

func (s *LinkStore) readLatest(limit int) ([]*models.Link, error) {
	rows, err := s.DB.Query(
		`SELECT * FROM links 
      ORDER BY published_at DESC, created_at DESC 
      LIMIT $1`,
		limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	data := []*models.Link{}
	for rows.Next() {
		d := &models.Link{}
		err = rows.Scan(
			&d.Id,
			&d.Url,
			&d.Title,
			&d.Description,
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

func (s *LinkStore) createOne(
	url string,
	title string,
	description string,
	significance string,
	publishedAt time.Time,
	tags []string,
) (int64, error) {
	result, err := s.DB.Exec(
		`INSERT INTO links (
      url, title, description, significance, published_at, tags
    ) VALUES (
      $1, $2, $3, $4, $5, $6
    )`,
		url, title, description, significance, publishedAt, tags)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *LinkStore) readOne(id int64) (*models.Link, error) {
	d := &models.Link{}

	err := s.DB.QueryRow(
		`SELECT * FROM links WHERE id = $1`, id).Scan(
		&d.Id,
		&d.Url,
		&d.Title,
		&d.Description,
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

func (s *LinkStore) updateOne(
	id int64,
	url string,
	title string,
	description string,
	significance int,
	publishedAt time.Time,
	tags []string,
) (int64, error) {
	result, err := s.DB.Exec(
		`UPDATE links SET (
      url = $2
      title = $3,
      description = $4,
      significance = $5,
      published_at = $6,
      tags = $7,
      updated_at = NOW()
    ) WHERE id = $1`,
		id, url, title, description, significance, publishedAt, tags)
	if err != nil {
		return 0, err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return n, nil
}

func (s *LinkStore) deleteOne(id int64) (int64, error) {
	result, err := s.DB.Exec(`DELETE FROM links WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return n, nil
}
