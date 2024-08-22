package models

import (
	"time"
)

type Note struct {
	Id           int64      `db:"id"`
	Title        string     `db:"title"`
	Body         *string    `db:"body"`
	Significance int        `db:"significance"`
	PublishedAt  *time.Time `db:"published_at"`
	Tags         []*string  `db:"tags"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`
}

type NoteRequest struct {
	Id           *int64     `json:"id"`
	Title        string     `json:"title"`
	Body         *string    `json:"body"`
	Significance int        `json:"significance"`
	PublishedAt  *time.Time `json:"published_at"`
	Tags         []*string  `json:"tags"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}
