package stores

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zettadam/adamz-api-go/internal/models"
)

type LinkStore struct {
	DB *pgxpool.Pool
}

func (s *LinkStore) ReadLatest(limit int) ([]models.Link, error) {
	rows, err := s.DB.Query(
		context.Background(),
		`SELECT * FROM links 
      ORDER BY published_at DESC, created_at DESC 
      LIMIT $1`,
		limit)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Link])
}

func (s *LinkStore) CreateOne(
	url string,
	title string,
	description string,
	significance string,
	publishedAt time.Time,
	tags []string,
) (int64, error) {
	var id int64 = 0
	err := s.DB.QueryRow(
		context.Background(),
		`INSERT INTO links (
      url, title, description, significance, published_at, tags
    ) VALUES (
      $1, $2, $3, $4, $5, $6
    ) RETURNING id`,
		url, title, description, significance, publishedAt, tags,
	).Scan(&id)
	return id, err
}

func (s *LinkStore) ReadOne(id int64) (models.Link, error) {
	rows, _ := s.DB.Query(
		context.Background(),
		`SELECT * FROM links WHERE id = $1`,
		id)
	return pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.Link])
}

func (s *LinkStore) UpdateOne(
	id int64,
	url string,
	title string,
	description string,
	significance int,
	publishedAt time.Time,
	tags []string,
) (int64, error) {
	result, err := s.DB.Exec(
		context.Background(),
		`UPDATE links SET (
      url = $2
      title = $3,
      description = $4,
      significance = $5,
      published_at = $6,
      tags = $7,
      updated_at = NOW()
    ) WHERE id = $1`,
		id, url, title, description, significance, publishedAt, tags,
	)
	return result.RowsAffected(), err
}

func (s *LinkStore) DeleteOne(id int64) (int64, error) {
	result, err := s.DB.Exec(
		context.Background(),
		`DELETE FROM links WHERE id = $1`,
		id)
	return result.RowsAffected(), err
}
