package stores

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zettadam/adamz-api-go/internal/models"
)

type PostStore struct {
	DB *pgxpool.Pool
}

func (s *PostStore) ReadLatest(limit int) ([]models.Post, error) {
	rows, err := s.DB.Query(
		context.Background(),
		`SELECT * FROM posts
      ORDER BY created_at DESC
      LIMIT $1`,
		limit)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Post])
}

func (s *PostStore) CreateOne(
	title string,
	slug string,
	abstract string,
	body string,
	significance int,
	publishedAt int,
	tags []string,
) (models.Post, error) {
	result, err := s.DB.Query(
		context.Background(),
		`INSERT INTO posts (
      title, slug, abstract, body, significance, published_at, tags
    ) VALUES (
      $1, $2, $3, $4, $5, $6, $7
    ) RETURNING *`,
		title, slug, abstract, body, significance, publishedAt, tags,
	)
	if err != nil {
		return models.Post{}, err
	}
	return pgx.CollectOneRow(result, pgx.RowToStructByPos[models.Post])
}

func (s *PostStore) ReadOne(id int64) (models.Post, error) {
	rows, err := s.DB.Query(
		context.Background(),
		`SELECT * FROM posts WHERE id = $1`,
		id)
	if err != nil {
		return models.Post{}, err
	}
	return pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.Post])
}

func (s *PostStore) UpdateOne(
	id int64,
	title string,
	slug string,
	abstract string,
	body string,
	significance int,
	publishedAt int,
	tags []string,
) (models.Post, error) {
	result, err := s.DB.Query(
		context.Background(),
		`UPDATE posts SET (
      title = $2,
      slug = $3,
      abstract = $4,
      body = $5,
      significance = $6,
      published_at = $7,
      tags = $8,
      updated_at = NOW()
    ) WHERE id = $1
    RETURNING *`,
		id, title, slug, abstract, body, significance, publishedAt, tags,
	)
	if err != nil {
		return models.Post{}, err
	}
	return pgx.CollectOneRow(result, pgx.RowToStructByPos[models.Post])
}

func (s *PostStore) DeleteOne(id int64) (int64, error) {
	result, err := s.DB.Exec(
		context.Background(),
		`DELETE FROM posts WHERE id = $1`,
		id)
	return result.RowsAffected(), err
}
