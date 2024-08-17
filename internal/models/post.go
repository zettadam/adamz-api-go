package models

import (
	"database/sql"
	"errors"
	"time"
)

type Post struct {
	Id           int64          `json:"id"`
	Title        string         `json:"title"`
	Slug         string         `json:"slug"`
	Abstract     sql.NullString `json:"abstract"`
	Body         sql.NullString `json:"body"`
	Significance int            `json:"significance"`
	Tags         sql.Null[any]  `json:"tags"`
	PublishedAt  sql.NullTime   `json:"published_at"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
}

type PostModel struct {
	DB *sql.DB
}

func (m *PostModel) readLatest(limit int) ([]*Post, error) {
	stmt := `SELECT * FROM posts
    ORDER BY published_at DESC, created_at DESC
    LIMIT $1`

	rows, err := m.DB.Query(stmt, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	data := []*Post{}
	for rows.Next() {
		d := &Post{}
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

func (m *PostModel) createOne(
	title string,
	slug string,
	abstract string,
	body string,
	significance int,
	publishedAt int,
	tags []string,
) (int64, error) {
	stmt := `INSERT INTO posts (
    title, slug, abstract, body, significance, published_at, tags
  ) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	result, err := m.DB.Exec(
		stmt,
		title, slug, body, publishedAt)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *PostModel) readOne(id int) (*Post, error) {
	stmt := `SELECT * FROM posts WHERE id = $1`
	d := &Post{}

	err := m.DB.QueryRow(stmt, id).Scan(
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
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return d, nil
}

func (m *PostModel) updateOne(
	id int,
	title string,
	slug string,
	abstract string,
	body string,
	significance int,
	publishedAt int,
	tags []string,
) (int64, error) {
	stmt := `UPDATE posts SET (
    title = $2,
    slug = $3,
    abstract = $4,
    body = $5,
    significance = $6,
    publishedAt = $7,
    tags = $8,
    updated_at = NOW()
  ) WHERE id = $1`

	result, err := m.DB.Exec(
		stmt,
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

func (m *PostModel) deleteOne(id int) (int64, error) {
	stmt := `DELETE FROM posts SET WHERE id = $1`

	result, err := m.DB.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return n, nil
}
