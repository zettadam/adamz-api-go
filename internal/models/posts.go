package models

import (
	"database/sql"
	"errors"
	"time"
)

type Post struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Body        string    `json:"body"`
	PublishedAt time.Time `json:"publishedAt"`
	CreatedAt   time.Time `json:"createdAt"`
}

type PostModel struct {
	DB *sql.DB
}

func (m *PostModel) ReadLatest(limit int) ([]*Post, error) {
	stmt := `SELECT * FROM posts 
    ORDER BY publishedAt DESC, createdAt DESC 
    LIMIT $1`

	rows, err := m.DB.Query(stmt, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := []*Post{}
	for rows.Next() {
		p := &Post{}
		err = rows.Scan(&p.Id, &p.Title, &p.Slug, &p.Body, &p.PublishedAt, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (m *PostModel) Create(title string, slug string, body string, publishedAt int) (int, error) {
	stmt := `INSERT INTO posts (
    title, slug, body, publishedAt
  ) VALUES ($1, $2, $3, $4)`

	result, err := m.DB.Exec(stmt, title, slug, body, publishedAt)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *PostModel) ReadOne(id int) (*Post, error) {
	stmt := `SELECT * FROM posts WHERE id = $1`
	p := &Post{}

	err := m.DB.QueryRow(stmt, id).Scan(
		&p.Id,
		&p.Title,
		&p.Slug,
		&p.Body,
		&p.PublishedAt,
		&p.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return p, nil
}

func (m *PostModel) Update(id int, title string, slug string, body string, publishedAt int) (int, error) {
	stmt := `UPDATE posts SET (
    title = $2,
    slug = $3,
    body = $4,
    publishedAt = $5
  ) WHERE id = $1`

	result, err := m.DB.Exec(stmt, id, title, slug, body, publishedAt)
	if err != nil {
		return 0, err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(n), nil
}

func (m *PostModel) Delete(id int) (int, error) {
	stmt := `DELETE FROM posts SET WHERE id = $1`

	result, err := m.DB.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(n), nil
}
