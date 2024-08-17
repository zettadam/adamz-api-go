package stores

import (
	"database/sql"
	"errors"
	"time"

	"github.com/zettadam/adamz-api-go/internal/models"
)

type CodeSnippetStore struct {
	DB *sql.DB
}

func (s *CodeSnippetStore) readLatest(limit int) ([]*models.CodeSnippet, error) {
	rows, err := s.DB.Query(
		`SELECT * FROM code_snippets
      ORDER BY published_at DESC, created_at DESC
      LIMIT $1`,
		limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	data := []*models.CodeSnippet{}
	for rows.Next() {
		d := &models.CodeSnippet{}
		err = rows.Scan(
			&d.Id,
			&d.Title,
			&d.Description,
			&d.Language,
			&d.Body,
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

func (s *CodeSnippetStore) createOne(
	title string,
	description string,
	language string,
	body string,
	publishedAt time.Time,
	tags []string,
) (int64, error) {
	result, err := s.DB.Exec(
		`INSERT INTO code_snippets (
      title, description, language, body, published_at, tags
    ) VALUES (
      $1, $2, $3, $4, $5, $6
    )`,
		title, description, language, body, publishedAt, tags)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *CodeSnippetStore) readOne(id int64) (*models.CodeSnippet, error) {
	d := &models.CodeSnippet{}

	err := s.DB.QueryRow(
		`SELECT * FROM links WHERE id = $1`, id).Scan(
		&d.Id,
		&d.Title,
		&d.Description,
		&d.Language,
		&d.Body,
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

func (s *CodeSnippetStore) updateOne(
	id int64,
	title string,
	description string,
	language string,
	body string,
	publishedAt string,
	tags []string,
) (int64, error) {
	result, err := s.DB.Exec(
		`UPDATE code_snippets SET (
      title = $2,
      description = $3,
      language = $4,
      body = $5
      published_at = $6,
      tags = $7,
      updated_at = NOW()
    ) WHERE id = $1`,
		id, title, description, language, body, publishedAt, tags)
	if err != nil {
		return 0, err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return n, nil
}

func (s *CodeSnippetStore) deleteOne(id int64) (int64, error) {
	return 0, nil
}
