package stores

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zettadam/adamz-api-go/internal/models"
)

type NoteStore struct {
	DB *pgxpool.Pool
}

func (s *NoteStore) ReadLatest(limit int) ([]models.Note, error) {
	rows, err := s.DB.Query(
		context.Background(),
		`SELECT * FROM notes
      ORDER BY published_at DESC, created_at DESC
      LIMIT $1`,
		limit)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Note])
}

func (s *NoteStore) CreateOne(
	title string,
	body string,
	significance int,
	publishedAt int,
	tags []string,
) (int64, error) {
	var id int64 = 0
	err := s.DB.QueryRow(
		context.Background(),
		`INSERT INTO notes (
      title, body, significance, published_at, tags
    ) VALUES (
      $1, $2, $3, $4, $5
    ) RETURNING id`,
		title, body, significance, publishedAt, tags,
	).Scan(&id)
	return id, err
}

func (s *NoteStore) ReadOne(id int64) (models.Note, error) {
	rows, _ := s.DB.Query(
		context.Background(),
		`SELECT * FROM notes WHERE id = $1`,
		id)
	return pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.Note])
}

func (s *NoteStore) UpdateOne(
	id int64,
	title string,
	body string,
	significance int,
	publishedAt time.Time,
	tags []string,
) (int64, error) {
	result, err := s.DB.Exec(
		context.Background(),
		`UPDATE notes SET (
      title = $2,
      body = $3,
      significance = $4,
      published_at = $5,
      tags = $6,
      updated_at = NOW()
    ) WHERE id = $1`,
		id, title, body, significance, publishedAt, tags)
	return result.RowsAffected(), err
}

func (s *NoteStore) DeleteOne(id int64) (int64, error) {
	result, err := s.DB.Exec(
		context.Background(),
		`DELETE FROM notes WHERE id = $1`,
		id)
	return result.RowsAffected(), err
}
