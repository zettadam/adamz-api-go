package models

import (
	"database/sql"
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
