package stores

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zettadam/adamz-api-go/internal/models"
)

type CodeSnippetStore struct {
	DB *pgxpool.Pool
}

func (s *CodeSnippetStore) ReadLatest(limit int) ([]models.CodeSnippet, error) {
	rows, err := s.DB.Query(
		context.Background(),
		`SELECT * FROM code_snippets
      ORDER BY published_at DESC, created_at DESC
      LIMIT $1`,
		limit)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.CodeSnippet])
}

func (s *CodeSnippetStore) CreateOne(
	title string,
	description string,
	language string,
	body string,
	publishedAt time.Time,
	tags []string,
) (int64, error) {
	var id int64 = 0
	err := s.DB.QueryRow(
		context.Background(),
		`INSERT INTO code_snippets (
      title, description, language, body, published_at, tags
    ) VALUES (
      $1, $2, $3, $4, $5, $6
    ) RETURNING id`,
		title, description, language, body, publishedAt, tags,
	).Scan(&id)
	return id, err
}

func (s *CodeSnippetStore) ReadOne(id int64) (*models.CodeSnippet, error) {
	rows, _ := s.DB.Query(
		context.Background(),
		`SELECT * FROM code_snippets WHERE id = $1`,
		id,
	)
	return pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[models.CodeSnippet])
}

func (s *CodeSnippetStore) UpdateOne(
	id int64,
	title string,
	description string,
	language string,
	body string,
	publishedAt string,
	tags []string,
) (int64, error) {
	result, err := s.DB.Exec(
		context.Background(),
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
	return result.RowsAffected(), err
}

func (s *CodeSnippetStore) DeleteOne(id int64) (int64, error) {
	result, err := s.DB.Exec(
		context.Background(),
		`DELETE FROM events WHERE id = $1`,
		id)
	return result.RowsAffected(), err
}
