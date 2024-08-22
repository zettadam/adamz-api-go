package stores

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zettadam/adamz-api-go/internal/types"
)

type LinkStore struct {
	DB *pgxpool.Pool
}

func (s *LinkStore) ReadLatest(limit int) ([]types.Link, error) {
	rows, err := s.DB.Query(
		context.Background(),
		`SELECT * FROM links 
      ORDER BY published_at DESC, created_at DESC 
      LIMIT $1`,
		limit)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByName[types.Link])
}

func (s *LinkStore) CreateOne(d types.LinkRequest) (types.Link, error) {
	result, err := s.DB.Query(
		context.Background(),
		`INSERT INTO links (
      url, title, description, significance, published_at, tags
    ) VALUES (
      $1, $2, $3, $4, $5, $6
    ) RETURNING *`,
		d.Url, d.Title, d.Description, d.Significance, d.PublishedAt, d.Tags,
	)
	if err != nil {
		return types.Link{}, err
	}
	return pgx.CollectOneRow(result, pgx.RowToStructByPos[types.Link])
}

func (s *LinkStore) ReadOne(id int64) (types.Link, error) {
	result, err := s.DB.Query(
		context.Background(),
		`SELECT * FROM links WHERE id = $1`,
		id)
	if err != nil {
		return types.Link{}, err
	}
	return pgx.CollectOneRow(result, pgx.RowToStructByPos[types.Link])
}

func (s *LinkStore) UpdateOne(
	id int64,
	d types.LinkRequest,
) (types.Link, error) {
	result, err := s.DB.Query(
		context.Background(),
		`UPDATE links SET (
      url = $2
      title = $3,
      description = $4,
      significance = $5,
      published_at = $6,
      tags = $7,
      updated_at = NOW()
    ) WHERE id = $1
    RETURNING *`,
		id,
		d.Url,
		d.Title,
		d.Description,
		d.Significance,
		d.PublishedAt,
		d.Tags,
	)
	if err != nil {
		return types.Link{}, err
	}
	return pgx.CollectOneRow(result, pgx.RowToStructByPos[types.Link])
}

func (s *LinkStore) DeleteOne(id int64) (int64, error) {
	result, err := s.DB.Exec(
		context.Background(),
		`DELETE FROM links WHERE id = $1`,
		id)
	return result.RowsAffected(), err
}
