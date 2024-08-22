package stores

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zettadam/adamz-api-go/internal/models"
)

type NoteStore struct {
	DB *pgxpool.Pool
}

func (s *NoteStore) ReadLatest(limit int) ([]models.Note, error) {
	result, err := s.DB.Query(
		context.Background(),
		`SELECT * FROM notes
      ORDER BY published_at DESC, created_at DESC
      LIMIT $1`,
		limit)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(result, pgx.RowToStructByName[models.Note])
}

func (s *NoteStore) CreateOne(d models.NoteRequest) (models.Note, error) {
	result, err := s.DB.Query(
		context.Background(),
		`INSERT INTO notes (
      title, body, significance, published_at, tags
    ) VALUES (
      $1, $2, $3, $4, $5
    ) RETURNING *`,
		d.Title, d.Body, d.Significance, d.PublishedAt, d.Tags,
	)
	if err != nil {
		return models.Note{}, err
	}

	return pgx.CollectOneRow(result, pgx.RowToStructByPos[models.Note])
}

func (s *NoteStore) ReadOne(id int64) (models.Note, error) {
	result, err := s.DB.Query(
		context.Background(),
		`SELECT * FROM notes WHERE id = $1`,
		id)
	if err != nil {
		return models.Note{}, err
	}
	return pgx.CollectOneRow(result, pgx.RowToStructByPos[models.Note])
}

func (s *NoteStore) UpdateOne(
	id int64,
	d models.NoteRequest,
) (models.Note, error) {
	result, err := s.DB.Query(
		context.Background(),
		`UPDATE notes SET (
      title = $2,
      body = $3,
      significance = $4,
      published_at = $5,
      tags = $6,
      updated_at = NOW()
    ) WHERE id = $1
    RETURNING *`,
		id,
		d.Title,
		d.Body,
		d.Significance,
		d.PublishedAt,
		d.Tags,
	)
	if err != nil {
		return models.Note{}, err
	}
	return pgx.CollectOneRow(result, pgx.RowToStructByPos[models.Note])
}

func (s *NoteStore) DeleteOne(id int64) (int64, error) {
	result, err := s.DB.Exec(
		context.Background(),
		`DELETE FROM notes WHERE id = $1`,
		id)
	return result.RowsAffected(), err
}
