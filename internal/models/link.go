package models

import (
	"database/sql"
	"time"
)

type Link struct {
	Id           int64          `json:"id"`
	Url          string         `json:"url"`
	Title        string         `json:"title"`
	Description  sql.NullString `json:"description"`
	Significance int            `json:"significance"`
	PublishedAt  sql.NullTime   `json:"published_at"`
	Tags         sql.Null[any]  `json:"tags"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
}
