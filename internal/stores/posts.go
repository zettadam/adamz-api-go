package stores

import (
	"database/sql"
	"errors"

	"github.com/zettadam/adamz-api-go/internal/models"
)

type PostStore struct {
	DB *sql.DB
}

func (s *PostStore) ReadLatest(limit int) ([]*models.Post, error) {
	rows, err := s.DB.Query(
		`SELECT * FROM posts
      ORDER BY created_at DESC
      LIMIT $1`,
		limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	data := []*models.Post{}
	for rows.Next() {
		d := &models.Post{}
		err = rows.Scan(
			&d.Id,
			&d.Title,
			&d.Slug,
			&d.Abstract,
			&d.Body,
			&d.Significance,
			&d.PublishedAt,
			&d.Tags,
			&d.CreatedAt,
			&d.UpdatedAt,
		)
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

func (s *PostStore) CreateOne(
	title string,
	slug string,
	abstract string,
	body string,
	significance int,
	publishedAt int,
	tags []string,
) (int64, error) {
	result, err := s.DB.Exec(
		`INSERT INTO posts (
      title, slug, abstract, body, significance, published_at, tags
    ) VALUES (
      $1, $2, $3, $4, $5, $6, $7
    )`,
		title, slug, abstract, body, significance, publishedAt, tags)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *PostStore) ReadOne(id int) (*models.Post, error) {
	d := &models.Post{}

	err := s.DB.QueryRow(
		`SELECT * FROM posts WHERE id = $1`, id).Scan(
		&d.Id,
		&d.Title,
		&d.Slug,
		&d.Abstract,
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

func (s *PostStore) UpdateOne(
	id int,
	title string,
	slug string,
	abstract string,
	body string,
	significance int,
	publishedAt int,
	tags []string,
) (int64, error) {
	result, err := s.DB.Exec(
		`UPDATE posts SET (
      title = $2,
      slug = $3,
      abstract = $4,
      body = $5,
      significance = $6,
      published_at = $7,
      tags = $8,
      updated_at = NOW()
    ) WHERE id = $1`,
		id, title, slug, abstract, body, significance, publishedAt, tags)
	if err != nil {
		return 0, err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return n, nil
}

func (s *PostStore) DeleteOne(id int) (int64, error) {
	result, err := s.DB.Exec(`DELETE FROM posts WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return n, nil
}
