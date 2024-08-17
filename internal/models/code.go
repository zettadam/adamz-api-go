package models

import (
	"database/sql"
	"time"
)

type CodeSnippet struct {
	Id          int64          `json:"id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Language    sql.NullString `json:"language"`
	Body        sql.NullString `json:"body"`
	PublishedAt sql.NullTime   `json:"published_at"`
	Tags        sql.Null[any]  `json:"tags"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"update_at"`
}
