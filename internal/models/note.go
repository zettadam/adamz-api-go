package models

import (
	"database/sql"
	"time"
)

type Note struct {
	Id           int64          `json:"id"`
	Title        string         `json:"title"`
	Body         sql.NullString `json:"body"`
	Significance int            `json:"significance"`
	PublishedAt  sql.NullTime   `json:"published_at"`
	Tags         sql.Null[any]  `json:"tags"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
}
