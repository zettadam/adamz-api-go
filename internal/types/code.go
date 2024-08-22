package types

import (
	"time"
)

type CodeSnippet struct {
	Id          int64      `db:"id"`
	Title       string     `db:"title"`
	Description *string    `db:"description"`
	Language    *string    `db:"language"`
	Body        *string    `db:"body"`
	PublishedAt *time.Time `db:"published_at"`
	Tags        []*string  `db:"tags"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"update_at"`
}

type CodeSnippetRequest struct {
	Id          *int64     `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Language    *string    `json:"language"`
	Body        *string    `json:"body"`
	PublishedAt *time.Time `json:"published_at"`
	Tags        []*string  `json:"tags"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"update_at"`
}
